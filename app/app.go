package app

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime/pprof"
	"runtime/trace"
	"sort"
	"strings"
	"time"

	"github.com/g3n/engine/audio/al"
	"github.com/g3n/engine/util"
	"github.com/kardianos/osext"

	"github.com/g3n/engine/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/util/logger"
	"github.com/g3n/engine/util/stats"
	"github.com/g3n/engine/window"
)

// App contains the application state
type App struct {
	*app.Application                  // Embedded standard application object
	log              *logger.Logger   // Application logger
	currentDemo      IDemo            // Current demo
	dirData          string           // Full path of the data directory
	scene            *core.Node       // Scene rendered
	demoScene        *core.Node       // Scene populated by individual demos
	ambLight         *light.Ambient   // Scene ambient light
	frameRater       *util.FrameRater // Render loop frame rater

	// GUI
	mainPanel  *gui.Panel
	demoPanel  *gui.Panel
	labelFPS   *gui.Label         // header FPS label
	treeTests  *gui.Tree          // tree with test names
	stats      *stats.Stats       // statistics object
	statsTable *stats.StatsTable  // statistics table panel
	control    *gui.ControlFolder // Pointer to gui control panel

	// Camera and orbit control
	camera *camera.Camera       // Camera
	orbit  *camera.OrbitControl // Orbit control
}

// IDemo is the interface that must be satisfied by all demos.
type IDemo interface {
	Start(*App)                 // Called once at the start of the demo
	Update(*App, time.Duration) // Called every frame
	Cleanup(*App)               // Called once at the end of the demo
}

// DemoMap maps the demo name string to its object
// Individual demos sets the keys of this map
var DemoMap = map[string]IDemo{}

// Command line options
var (
	// TODO uncomment and implement usage of the following flags
	//oFullScreen   = flag.Bool("fullscreen", false, "Starts application with full screen")
	//oSwapInterval = flag.Int("swapinterval", -1, "Sets the swap buffers interval to this value")
	oHideFPS     = flag.Bool("hidefps", false, "Do now show calculated FPS in the GUI")
	oUpdateFPS   = flag.Uint("updatefps", 1000, "Time interval in milliseconds to update the FPS in the GUI")
	oTargetFPS   = flag.Uint("targetfps", 60, "Sets the frame rate in frames per second")
	oNoglErrors  = flag.Bool("noglerrors", false, "Do not check OpenGL errors at each call (may increase FPS)")
	oCpuProfile  = flag.String("cpuprofile", "", "Activate cpu profiling writing profile to the specified file")
	oExecTrace   = flag.String("exectrace", "", "Activate execution tracer writing data to the specified file")
	oNogui       = flag.Bool("nogui", false, "Do not show the GUI, only the specified demo")
	oLogs        = flag.String("logs", "", "Set log levels for packages. Ex: gui:debug,gls:info")
	oStats       = flag.Bool("stats", false, "Shows statistics control panel in the GUI")
	oRenderStats = flag.Bool("renderstats", false, "Shows gui renderer statistics in the console")
)

// usage shows the usage of command line flags
func usage() {
	fmt.Fprintf(os.Stderr, "%s v%d.%d\n", progName, vmajor, vminor)
	fmt.Fprintf(os.Stderr, "usage: %s [options] [<test>] \n", execName)
	flag.PrintDefaults()
	os.Exit(2)
}

const (
	progName = "G3N Demo" // TODO set title (create pair of files for build tags)
	execName = "g3nd"
	vmajor   = 0
	vminor   = 6
)

// Create creates the G3ND application using the specified map of demos
func Create() *App {

	a := new(App)
	a.Application = app.App()

	// Creates application logger
	a.log = logger.New("G3ND", nil)
	a.log.AddWriter(logger.NewConsole(false))
	a.log.SetFormat(logger.FTIME | logger.FMICROS)
	a.log.SetLevel(logger.DEBUG)

	a.log.Info("%s v%d.%d starting", progName, vmajor, vminor)

	a.stats = stats.NewStats(a.Gls())

	// Log OpenGL version
	glVersion := a.Gls().GetString(gls.VERSION)
	a.log.Info("OpenGL version: %s", glVersion)

	// Set OpenGL error checking based on flag
	a.Gls().SetCheckErrors(!*oNoglErrors)

	// Create scenes
	a.demoScene = core.NewNode() // demoScene will be cleared before a new demo is started
	a.scene = core.NewNode()
	a.scene.Add(a.demoScene)

	// Create camera and orbit control
	width, height := a.GetSize()
	aspect := float32(width) / float32(height)
	a.camera = camera.New(aspect)
	a.scene.Add(a.camera) // Add camera to scene (important for audio demos)
	a.orbit = camera.NewOrbitControl(a.camera)

	// Create and add ambient light to scene
	a.ambLight = light.NewAmbient(&math32.Color{1.0, 1.0, 1.0}, 0.5)
	a.scene.Add(a.ambLight)

	// Create frame rater
	a.frameRater = util.NewFrameRater(*oTargetFPS)

	flag.Usage = usage // Sets the application usage
	flag.Parse()       // Parse command line flags

	// Apply log levels to engine package loggers
	if *oLogs != "" {
		logs := strings.Split(*oLogs, ",")
		for i := 0; i < len(logs); i++ {
			parts := strings.Split(logs[i], ":")
			if len(parts) != 2 {
				a.log.Error("Invalid logs level string")
				continue
			}
			pack := strings.ToUpper(parts[0])
			level := strings.ToUpper(parts[1])
			path := "G3N/" + pack
			packlog := logger.Find(path)
			if packlog == nil {
				a.log.Error("No logger for package:%s", pack)
				continue
			}
			err := packlog.SetLevelByName(level)
			if err != nil {
				a.log.Error("%s", err)
			}
			a.log.Info("Set log level:%s for package:%s", level, pack)
		}
	}

	// Check for data directory and abort if not found
	a.dirData = a.checkDirData("data")
	a.log.Info("Using data directory:%s", a.dirData)

	// Create demoPanel to house GUI elements created by the demos
	a.demoPanel = gui.NewPanel(0, 0)
	a.demoPanel.SetColor4(&gui.StyleDefault().Scroller.BgColor)
	a.demoPanel.SetLayoutParams(&gui.DockLayoutParams{Edge: gui.DockCenter})

	// Build user interface
	if *oNogui {
		a.scene.Add(a.demoPanel)
	} else {
		a.buildGui(DemoMap)
	}

	// Sets the default window resize event handler
	a.Subscribe(window.OnWindowSize, func(evname string, ev interface{}) { a.OnWindowResize() })
	a.OnWindowResize()

	// Subscribe to key events
	a.Subscribe(window.OnKeyDown, func(evname string, ev interface{}) {
		kev := ev.(*window.KeyEvent)
		if kev.Key == window.KeyEscape { // ESC terminates the program
			a.Exit()
		} else if kev.Key == window.KeyF11 { // F11 toggles full screen
			//a.Window().SetFullScreen(!a.Window().FullScreen()) // TODO
		} else if kev.Key == window.KeyS && kev.Mods == window.ModAlt { // Ctr-S prints statistics in the console
			a.logStats()
		}
	})

	// Setup scene
	a.setupScene()

	// If name of test supplied in the command line
	// set it as the current test and initialize it.
	if len(flag.Args()) > 0 {
		tname := flag.Args()[0]
		a.log.Info("ARGS")
		test, ok := DemoMap[tname]
		if ok {
			a.log.Info("Start")
			a.currentDemo = test
			a.currentDemo.Start(a)
		}
		if a.currentDemo == nil {
			a.log.Error("Invalid demo name")
			usage()
			return nil
		}
	}

	return a
}

// checkDirData try to find and return the complete data directory path.
// Aborts if not found
func (a *App) checkDirData(dirDataName string) string {

	// Check first if data directory is in the current directory
	if _, err := os.Stat(dirDataName); err == nil {
		dirData, err := filepath.Abs(dirDataName)
		if err != nil {
			panic(err)
		}
		return dirData
	}

	// Get the executable path
	execPath, err := osext.Executable()
	if err != nil {
		panic(err)
	}

	// Checks if data directory is in the executable directory
	execDir := filepath.Dir(execPath)
	path := filepath.Join(execDir, dirDataName)
	if _, err := os.Stat(path); err == nil {
		return path
	}

	// Assumes the executable is in $GOPATH/bin
	goPath := filepath.Dir(execDir)
	path = filepath.Join(goPath, "src", "github.com", "g3n", "g3nd", dirDataName)
	// Checks data path
	if _, err := os.Stat(path); err == nil {
		return path
	}

	// If the data directory hasn't been found, manually scan the $GOPATH directories
	rawPaths := os.Getenv("GOPATH")
	paths := strings.Split(rawPaths, ":")
	for _, j := range paths {
		// Checks data path
		path = filepath.Join(j, "src", "github.com", "g3n", "g3nd", dirDataName)
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	// Show error message and aborts
	a.log.Fatal("Data directory NOT FOUND")
	return ""
}

// logStats generate log with current statistics
func (a *App) logStats() {

	const statsFormat = `
         Shaders: %d
            Vaos: %d
         Buffers: %d
        Textures: %d
  Uniforms/frame: %d
Draw calls/frame: %d
 CGO calls/frame: %d
`
	a.log.Info(statsFormat,
		a.stats.Glstats.Shaders,
		a.stats.Glstats.Vaos,
		a.stats.Glstats.Buffers,
		a.stats.Glstats.Textures,
		a.stats.Unisets,
		a.stats.Drawcalls,
		a.stats.Cgocalls,
	)
}

// buildGui builds the tester GUI
func (a *App) buildGui(demoMap map[string]IDemo) {

	// Create dock layout for the tester root panel
	dl := gui.NewDockLayout()
	width, height := a.GetSize()
	a.mainPanel = gui.NewPanel(float32(width), float32(height))
	a.mainPanel.SetRenderable(false)
	a.mainPanel.SetEnabled(false)
	a.mainPanel.SetLayout(dl)
	a.scene.Add(a.mainPanel)
	gui.Manager().Set(a.mainPanel)

	// Add transparent panel at the center to contain demos
	a.mainPanel.Add(a.demoPanel)

	// Adds header after the gui central panel to ensure that the control folder
	// stays over the gui panel when opened.
	headerColor := math32.Color4{13.0 / 256.0, 41.0 / 256.0, 62.0 / 256.0, 1}
	lightTextColor := math32.Color4{0.8, 0.8, 0.8, 1}
	header := gui.NewPanel(600, 40)
	header.SetBorders(0, 0, 1, 0)
	header.SetPaddings(4, 4, 4, 4)
	header.SetColor4(&headerColor)
	header.SetLayoutParams(&gui.DockLayoutParams{Edge: gui.DockTop})

	// Horizontal box layout for the header
	hbox := gui.NewHBoxLayout()
	header.SetLayout(hbox)
	a.mainPanel.Add(header)

	// Add an optional image to header
	logo, err := gui.NewImage(a.dirData + "/images/g3n_logo_32.png")
	if err == nil {
		logo.SetContentAspectWidth(32)
		header.Add(logo)
	}

	// Header title
	const fontSize = 20
	title := gui.NewLabel(" ")
	title.SetFontSize(fontSize)
	title.SetLayoutParams(&gui.HBoxLayoutParams{AlignV: gui.AlignCenter})
	title.SetText(fmt.Sprintf("%s v%d.%d", progName, vmajor, vminor))
	title.SetColor4(&lightTextColor)
	header.Add(title)
	// FPS
	if !*oHideFPS {
		l1 := gui.NewLabel(" ")
		l1.SetFontSize(fontSize)
		l1.SetLayoutParams(&gui.HBoxLayoutParams{AlignV: gui.AlignCenter})
		l1.SetText("  FPS: ")
		l1.SetColor4(&lightTextColor)
		header.Add(l1)
		// FPS value
		a.labelFPS = gui.NewLabel(" ")
		a.labelFPS.SetFontSize(fontSize)
		a.labelFPS.SetLayoutParams(&gui.HBoxLayoutParams{AlignV: gui.AlignCenter})
		a.labelFPS.SetColor4(&lightTextColor)
		header.Add(a.labelFPS)
	}

	// New styles for control folder
	styles := gui.StyleDefault().ControlFolder
	styles.Folder.Normal.BgColor = headerColor
	styles.Folder.Over.BgColor = headerColor
	styles.Folder.Normal.FgColor = lightTextColor
	styles.Folder.Over.FgColor = lightTextColor

	// Adds statistics table control folder if requested
	if *oStats {
		// Adds spacer to right justify the control folder in the header
		spacer := gui.NewPanel(0, 0)
		spacer.SetLayoutParams(&gui.HBoxLayoutParams{AlignV: gui.AlignBottom, Expand: 1.2})
		header.Add(spacer)

		// Creates control folder for statistics table
		statsControlFolder := gui.NewControlFolder("Stats", 100)
		statsControlFolder.SetLayoutParams(&gui.HBoxLayoutParams{AlignV: gui.AlignBottom})
		statsControlFolder.SetStyles(&styles)
		header.Add(statsControlFolder)

		// Adds stats table in the control folder
		a.statsTable = stats.NewStatsTable(220, 200, a.Gls())
		statsControlFolder.AddPanel(a.statsTable)
	}

	// Adds spacer to right justify the control folder in the header
	spacer := gui.NewPanel(0, 0)
	spacer.SetLayoutParams(&gui.HBoxLayoutParams{AlignV: gui.AlignBottom, Expand: 1})
	header.Add(spacer)

	// Adds control folder in the header
	a.control = gui.NewControlFolder("Controls", 100)
	a.control.SetLayoutParams(&gui.HBoxLayoutParams{AlignV: gui.AlignBottom})
	a.control.SetStyles(&styles)
	header.Add(a.control)

	// Test list
	a.treeTests = gui.NewTree(175, 0)

	// TODO This does not persist - have to change style / but better yet is to improve GUI so that individual style changes can be performed this way
	//a.treeTests.SetBorders(0, 1, 1, 1)

	a.treeTests.SetLayoutParams(&gui.DockLayoutParams{Edge: gui.DockLeft})
	// Sort test names
	tnames := []string{}
	nodes := make(map[string]*gui.TreeNode)
	for name := range demoMap {
		tnames = append(tnames, name)
	}
	sort.Strings(tnames)
	// Add items to the list
	for _, name := range tnames {
		parts := strings.Split(name, ".")
		if len(parts) > 1 {
			category := parts[0]
			node := nodes[category]
			if node == nil {
				node = a.treeTests.AddNode(category)
				nodes[category] = node
			}
			labelText := strings.Join(parts[1:], ".")
			item := gui.NewLabel(labelText)
			item.SetUserData(demoMap[name])
			node.Add(item)
		} else {
			item := gui.NewLabel(name)
			item.SetUserData(demoMap[name])
			a.treeTests.Add(item)
		}
	}
	a.treeTests.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		sel := a.treeTests.Selected()
		if sel == nil {
			return
		}
		label, ok := sel.(*gui.Label)
		if ok {
			a.setupScene()
			test := label.GetNode().UserData().(IDemo)
			test.Start(a)
			a.currentDemo = test
		}
	})
	a.mainPanel.Add(a.treeTests)
}

// setupScene resets the current scene for executing a new (or first) test
func (a *App) setupScene() {

	// If there was a previous demo running, execute its Cleanup() method
	if a.currentDemo != nil {
		a.currentDemo.Cleanup(a)
	}

	// Destroy all objects in demo scene and GUI
	a.demoScene.DisposeChildren(true)
	a.demoPanel.DisposeChildren(true)

	// By default set the demo panel as not renderable (so it doesn't show) and not enabled (so it doesn't capture events)
	a.demoPanel.SetRenderable(false)
	a.demoPanel.SetEnabled(false)

	// Clear subscriptions with ID (every subscribe called by demos should use the app address as ID so we can unsubscribe here)
	a.demoPanel.UnsubscribeAllID(a)
	a.UnsubscribeAllID(a)

	// Clear all custom cursors and reset current cursor
	a.DisposeAllCustomCursors()
	a.SetCursor(window.ArrowCursor)

	// Set default background color
	a.Gls().ClearColor(0.6, 0.6, 0.6, 1.0)

	// Reset renderer z-sorting flag
	a.Renderer().SetObjectSorting(true)

	// Reset ambient light
	a.ambLight.SetColor(&math32.Color{1.0, 1.0, 1.0})
	a.ambLight.SetIntensity(0.5)

	// Reset camera
	a.camera.SetPosition(0, 0, 5)
	a.camera.UpdateSize(5)
	a.camera.LookAt(&math32.Vector3{0, 0, 0}, &math32.Vector3{0, 1, 0})
	a.camera.SetProjection(camera.Perspective)
	a.orbit.Reset()

	// If audio active, resets global listener parameters
	al.Listener3f(al.Position, 0, 0, 0)
	al.Listener3f(al.Velocity, 0, 0, 0)
	al.Listenerfv(al.Orientation, []float32{0, 0, -1, 0, 1, 0})

	// If no gui control folder, nothing more to do
	if a.control == nil {
		return
	}

	// Remove all controls and adds default ones
	a.control.Clear()

	// Adds camera selection
	cb := a.control.AddCheckBox("Perspective camera").SetValue(true)
	cb.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		if cb.Value() {
			a.camera.SetProjection(camera.Perspective)
		} else {
			a.camera.SetProjection(camera.Orthographic)
		}
	})

	// Adds ambient light slider
	s1 := a.control.AddSlider("Ambient light:", 2.0, a.ambLight.Intensity())
	s1.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		a.ambLight.SetIntensity(s1.Value())
	})
}

// DemoPanel returns the current gui panel for demos to add elements to.
func (a *App) DemoPanel() *gui.Panel {

	return a.demoPanel
}

// DirData returns the base directory for data
func (a *App) DirData() string {

	return a.dirData
}

// ControlFolder returns the application control folder
func (a *App) ControlFolder() *gui.ControlFolder {

	return a.control
}

// AmbLight returns the default scene ambient light
func (a *App) AmbLight() *light.Ambient {

	return a.ambLight
}

// Log returns the application logger
func (a *App) Log() *logger.Logger {

	return a.log
}

// Scene returns the current application 3D scene
func (a *App) Scene() *core.Node {

	return a.demoScene
}

// Camera returns the current application camera
func (a *App) Camera() *camera.Camera {

	return a.camera
}

// Orbit returns the current camera orbit control
func (a *App) Orbit() *camera.OrbitControl {

	return a.orbit
}

// OnWindowResize is default handler for window resize events.
func (a *App) OnWindowResize() {

	// Get framebuffer size and set the viewport accordingly
	width, height := a.GetFramebufferSize()
	a.Gls().Viewport(0, 0, int32(width), int32(height))

	// Set camera aspect ratio
	a.camera.SetAspect(float32(width) / float32(height))

	if *oNogui {
		a.demoPanel.SetSize(float32(width), float32(height))
	} else {
		a.mainPanel.SetSize(float32(width), float32(height))
	}
}

// Run runs the application render loop
func (a *App) Run() {

	// Start profiling if requested
	if *oCpuProfile != "" {
		f, err := os.Create(*oCpuProfile)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		err = pprof.StartCPUProfile(f)
		if err != nil {
			panic(err)
		}
		defer pprof.StopCPUProfile()
		a.log.Info("Started writing CPU profile to: %s", *oCpuProfile)
	}

	// Start execution trace if requested
	if *oExecTrace != "" {
		f, err := os.Create(*oExecTrace)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		err = trace.Start(f)
		if err != nil {
			panic(err)
		}
		defer trace.Stop()
		a.log.Info("Started writing execution trace to: %s", *oExecTrace)
	}

	a.Application.Run(a.Update)
}

func (a *App) Update(rend *renderer.Renderer, deltaTime time.Duration) {

	// Start measuring this frame
	a.frameRater.Start()

	// Clear the color, depth, and stencil buffers
	a.Gls().Clear(gls.COLOR_BUFFER_BIT | gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT) // TODO maybe do inside renderer, and allow customization

	// Update the current running demo if any
	if a.currentDemo != nil {
		a.currentDemo.Update(a, deltaTime)
	}

	// Render scene
	err := rend.Render(a.scene, a.camera)
	if err != nil {
		panic(err)
	}

	// Update GUI timers
	gui.Manager().TimerManager.ProcessTimers()

	// Update statistics
	if a.stats.Update(time.Second) {
		if a.statsTable != nil {
			a.statsTable.Update(a.stats)
		}
	}

	// Update render stats
	if *oRenderStats {
		stats := a.Renderer().Stats()
		if stats.Panels > 0 {
			a.log.Debug("render stats:%+v", stats)
		}
	}

	// Control and update FPS
	a.frameRater.Wait()
	a.updateFPS()
}

// UpdateFPS updates the fps value in the window title or header label
func (a *App) updateFPS() {

	if *oHideFPS {
		return
	}

	// Get the FPS and potential FPS from the frameRater
	fps, pfps, ok := a.frameRater.FPS(time.Duration(*oUpdateFPS) * time.Millisecond)
	if !ok {
		return
	}

	// Show the FPS in the header label
	a.labelFPS.SetText(fmt.Sprintf("%3.1f / %3.1f", fps, pfps))
}

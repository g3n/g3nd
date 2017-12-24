package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"sort"
	"strings"
	"time"

	"github.com/g3n/engine/audio/al"
	"github.com/g3n/engine/audio/ov"
	"github.com/g3n/engine/audio/vorbis"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/camera/control"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/util/logger"
	"github.com/g3n/engine/util/stats"
	"github.com/g3n/engine/window"
	"github.com/kardianos/osext"
)

const (
	ProgName = "G3N Demo"
	ExecName = "g3nd"
	Vmajor   = 0
	Vminor   = 4
)

// Package logger
var log *logger.Logger

// Data directory base name
const dirDataName = "data"

// Context structure passed to all tests
type Context struct {
	GS          *gls.GLS              // OpenGL state
	Win         window.IWindow        // Window
	Renderer    *renderer.Renderer    // pointer to renderer object
	CamPersp    *camera.Perspective   // pointer to perspective camera
	CamOrtho    *camera.Orthographic  // pointer to orthographic camera
	Camera      camera.ICamera        // current camera
	Orbit       *control.OrbitControl // pointer to orbit camera controller
	Gui         *gui.Panel            // GUI panel container for GUI tests
	Control     *gui.ControlFolder    // Pointer to gui control panel
	Scene       *core.Node            // Node container for 3D tests
	AmbLight    *light.Ambient        // pointer to ambient light
	DirData     string                // directory of test data files
	Time        time.Time             // current time at the start of the frame
	TimeDelta   time.Duration         // time delta from previous frame
	Audio       bool                  // Audio available
	AudioEFX    bool                  // Audio effects available
	Vorbis      bool                  // Vorbis decoder available
	AudioDev    *al.Device            // Audio player device
	CapDev      *al.Device            // Audio capture device
	root        *gui.Root             // GUI root container
	currentTest ITest                 // current test object
	labelFPS    *gui.Label            // header FPS label
	treeTests   *gui.Tree             // tree with test names
	frameRater  *FrameRater           // frame rate controller
	stats       *stats.Stats          // statistics object
	statsTable  *stats.StatsTable     // statistics table panel
}

// ITest is the interface that must be satisfied for all test objects
type ITest interface {
	Initialize(*Context)
	Render(*Context)
}

// TestMap maps the test name string to its object
var TestMap = map[string]ITest{}

// Command line options
var (
	oVersion    = flag.Bool("version", false, "Show version and exits")
	oWidth      = flag.Int("width", 1000, "Initial window width in pixels")
	oHeight     = flag.Int("height", 800, "Initial window height in pixels")
	oFull       = flag.Bool("full", false, "Full screen on primary monitor")
	oNogui      = flag.Bool("nogui", false, "Do not show the GUI, only the specified demo")
	oHideFPS    = flag.Bool("hidefps", false, "Do now show calculated FPS in the GUI")
	oUpdateFPS  = flag.Uint("updatefps", 1000, "Time interval in milliseconds to update the FPS in the GUI")
	oFPS        = flag.Uint("fps", 60, "Sets the frame rate in frames per second")
	oInterval   = flag.Int("interval", 0, "Sets the swap buffers interval to this value")
	oLogColor   = flag.Bool("logcolors", false, "Colored logs")
	oLogs       = flag.String("logs", "", "Set log levels for packages. Ex: gui:debug,gls:info")
	oNoGlErrors = flag.Bool("noglerrors", false, "Do not check OpenGL errors at each call (may increase FPS)")
	oProfile    = flag.String("profile", "", "Activate cpu profiling writing profile to the specified file")
	oStats      = flag.Bool("stats", false, "Shows statistics control panel in the GUI")
)

func main() {

	// Parse command line parameters
	flag.Usage = usage
	flag.Parse()

	// If requested, print version and exits
	if *oVersion == true {
		fmt.Fprintf(os.Stderr, "%s v%d.%d\n", ProgName, Vmajor, Vminor)
		return
	}

	// Creates independent logger for the application
	log = logger.New("G3ND", nil)
	log.AddWriter(logger.NewConsole(*oLogColor))
	log.SetFormat(logger.FTIME | logger.FMICROS)
	log.SetLevel(logger.DEBUG)
	log.Info("%s v%d.%d starting", ProgName, Vmajor, Vminor)

	// Apply log levels to engine package loggers
	if *oLogs != "" {
		logs := strings.Split(*oLogs, ",")
		for i := 0; i < len(logs); i++ {
			parts := strings.Split(logs[i], ":")
			if len(parts) != 2 {
				log.Error("Invalid logs level string")
				continue
			}
			pack := strings.ToUpper(parts[0])
			level := strings.ToUpper(parts[1])
			path := "G3N/" + pack
			packlog := logger.Find(path)
			if packlog == nil {
				log.Error("No logger for package:%s", pack)
				continue
			}
			err := packlog.SetLevelByName(level)
			if err != nil {
				log.Error("%s", err)
			}
			log.Info("Set log level:%s for package:%s", level, pack)
		}
	}

	// Check for data directory and aborts if not found
	dirData := checkDirData()
	log.Info("Using data directory:%s", dirData)

	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()

	// Creates window and sets it as the current context
	win, err := window.New("glfw", *oWidth, *oHeight, "G3ND", *oFull)
	if err != nil {
		panic(err)
	}

	// Create OpenGL state
	gs, err := gls.New()
	if err != nil {
		panic(err)
	}
	glVersion := gs.GetString(gls.VERSION)
	log.Info("OpenGL version: %s", glVersion)
	gs.SetCheckErrors(!*oNoGlErrors)

	// Set swap buffers interval
	// If users passes -1, the system swap interval is unchanged
	if *oInterval >= 0 {
		win.SwapInterval(*oInterval)
	}

	// Starts building context which is passed to all tests
	var ctx Context
	ctx.GS = gs
	ctx.Win = win
	ctx.Time = time.Now()
	ctx.DirData = dirData
	ctx.frameRater = NewFrameRater(*oFPS)
	ctx.stats = stats.NewStats(gs)

	// Try to load audio libraries and sets its availability in the context
	loadAudioLibs(&ctx)

	// Creates renderer
	ctx.Renderer = renderer.NewRenderer(gs)
	err = ctx.Renderer.AddDefaultShaders()
	if err != nil {
		log.Fatal("AddDefaultShaders:%s", err)
	}

	// Creates scene for 3D objects
	ctx.Scene = core.NewNode()

	// Creates root panel for GUI
	ctx.root = gui.NewRoot(gs, win)
	if *oNogui {
		ctx.Gui = ctx.root.GetPanel()
	} else {
		buildGui(&ctx)
		ctx.Renderer.SetGuiPanel3D(ctx.Gui)
	}
	ctx.Renderer.SetGui(ctx.root)

	// Setup scene
	setupScene(&ctx)
	winResizeEvent(&ctx)

	// If name of test supplied in the command line
	// sets it as the current test and initialize it.
	if len(flag.Args()) > 0 {
		tname := flag.Args()[0]
		for name, test := range TestMap {
			if name == tname {
				ctx.currentTest = test
				ctx.currentTest.Initialize(&ctx)
				break
			}
		}
		if ctx.currentTest == nil {
			log.Error("INVALID TEST NAME")
			usage()
			return
		}
	}

	// Start profiling if requested
	if *oProfile != "" {
		f, err := os.Create(*oProfile)
		if err != nil {
			log.Fatal("Error creating profile file:%s", err)
		}
		err = pprof.StartCPUProfile(f)
		if err != nil {
			log.Fatal("%s", err)
		}
		log.Info("Started writing CPU profile to:%s", *oProfile)
		defer pprof.StopCPUProfile()
	}

	// Render loop
	for !win.ShouldClose() {
		// Starts measuring this frame
		ctx.frameRater.Start()

		// Updates time and time delta in context
		now := time.Now()
		ctx.TimeDelta = now.Sub(ctx.Time)
		ctx.Time = now

		// Process root panel timers
		ctx.root.TimerManager.ProcessTimers()

		// If current test active, render test scene
		if ctx.currentTest != nil {
			ctx.currentTest.Render(&ctx)
			ctx.Renderer.SetScene(ctx.Scene)
		} else {
			ctx.Renderer.SetScene(nil)
		}

		// Render Scene and/or Gui
		err := ctx.Renderer.Render(ctx.Camera)
		if err != nil {
			log.Fatal("Render error: %s\n", err)
		}

		// Update statistics
		if ctx.stats.Update(time.Second) {
			// Update statistics table
			if ctx.statsTable != nil {
				ctx.statsTable.Update(ctx.stats)
			}
		}

		// Poll input events and process them
		win.PollEvents()

		// Swap window framebuffers if necessary
		if ctx.Renderer.NeedSwap() {
			win.SwapBuffers()
			//log.Error("swap buffers")
		}
		// Controls the frame rate and updates the FPS for the user
		ctx.frameRater.Wait()
		updateFPS(&ctx)
	}
}

// buildGui builds the tester GUI
func buildGui(ctx *Context) {

	// Create dock layout for the tester root panel
	dl := gui.NewDockLayout()
	ctx.root.SetLayout(dl)

	// Add transparent panel at the center to contain demos
	ctx.Gui = gui.NewPanel(0, 0)
	ctx.Gui.SetRenderable(false)
	ctx.Gui.SetLayoutParams(&gui.DockLayoutParams{Edge: gui.DockCenter})
	ctx.root.Add(ctx.Gui)

	// Adds header after the gui central panel to ensure that the control folder
	// stays over the gui panel when opened.
	headerColor := math32.Color{0, 0.15, 0.3}
	lightTextColor := math32.Color{0.8, 0.8, 0.8}
	header := gui.NewPanel(600, 40)
	header.SetBorders(0, 0, 0, 0)
	header.SetPaddings(4, 4, 4, 4)
	header.SetColor(&headerColor)
	header.SetLayoutParams(&gui.DockLayoutParams{Edge: gui.DockTop})

	// Horizontal box layout for the header
	hbox := gui.NewHBoxLayout()
	header.SetLayout(hbox)
	ctx.root.Add(header)

	// Add an optional image to header
	logo, err := gui.NewImage(ctx.DirData + "/images/g3n_logo_32.png")
	if err == nil {
		logo.SetContentAspectWidth(32)
		header.Add(logo)
	}

	// Header title
	const fontSize = 20
	title := gui.NewLabel(" ")
	title.SetFontSize(fontSize)
	title.SetLayoutParams(&gui.HBoxLayoutParams{AlignV: gui.AlignCenter})
	title.SetText(fmt.Sprintf("%s v%d.%d", ProgName, Vmajor, Vminor))
	title.SetColor(&lightTextColor)
	header.Add(title)
	// FPS
	if !*oHideFPS {
		l1 := gui.NewLabel(" ")
		l1.SetFontSize(fontSize)
		l1.SetLayoutParams(&gui.HBoxLayoutParams{AlignV: gui.AlignCenter})
		l1.SetText("  FPS: ")
		l1.SetColor(&lightTextColor)
		header.Add(l1)
		// FPS value
		ctx.labelFPS = gui.NewLabel(" ")
		ctx.labelFPS.SetFontSize(fontSize)
		ctx.labelFPS.SetLayoutParams(&gui.HBoxLayoutParams{AlignV: gui.AlignCenter})
		ctx.labelFPS.SetColor(&lightTextColor)
		header.Add(ctx.labelFPS)
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
		ctx.statsTable = stats.NewStatsTable(220, 200, ctx.GS)
		statsControlFolder.AddPanel(ctx.statsTable)
	}

	// Adds spacer to right justify the control folder in the header
	spacer := gui.NewPanel(0, 0)
	spacer.SetLayoutParams(&gui.HBoxLayoutParams{AlignV: gui.AlignBottom, Expand: 1})
	header.Add(spacer)

	// Adds control folder in the header
	ctx.Control = gui.NewControlFolder("Controls", 100)
	ctx.Control.SetLayoutParams(&gui.HBoxLayoutParams{AlignV: gui.AlignBottom})
	ctx.Control.SetStyles(&styles)
	header.Add(ctx.Control)

	// Test list
	ctx.treeTests = gui.NewTree(150, 0)
	ctx.treeTests.SetLayoutParams(&gui.DockLayoutParams{Edge: gui.DockLeft})
	// Sort test names
	tnames := []string{}
	nodes := make(map[string]*gui.TreeNode)
	for name, _ := range TestMap {
		tnames = append(tnames, name)
	}
	sort.Strings(tnames)
	// Add items to the list
	for _, name := range tnames {
		parts := strings.Split(name, ".")
		if len(parts) > 1 {
			category := parts[0]
			// Do not include "audio" demos if vorbis not supported
			if category == "audio" && !ctx.Vorbis {
				continue
			}
			node := nodes[category]
			if node == nil {
				node = ctx.treeTests.AddNode(category)
				nodes[category] = node
			}
			labelText := strings.Join(parts[1:], ".")
			item := gui.NewLabel(labelText)
			item.SetUserData(TestMap[name])
			node.Add(item)
		} else {
			item := gui.NewLabel(name)
			item.SetUserData(TestMap[name])
			ctx.treeTests.Add(item)
		}
	}
	ctx.treeTests.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		sel := ctx.treeTests.Selected()
		if sel == nil {
			return
		}
		label, ok := sel.(*gui.Label)
		if ok {
			setupScene(ctx)
			test := label.GetNode().UserData().(ITest)
			test.Initialize(ctx)
			ctx.currentTest = test
		}
	})
	ctx.root.Add(ctx.treeTests)

}

// FrameRater implements a frame rate controller
type FrameRater struct {
	targetFPS      uint          // desired number of frames per second
	targetDuration time.Duration // calculated desired duration of frame
	frameStart     time.Time     // start time of last frame
	frameTimes     time.Duration // accumulated frame times for potential FPS calculation
	frameCount     uint          // accumulated number of frames for FPS calculation
	lastUpdate     time.Time     // time of last FPS calculation update
	timer          *time.Timer   // timer for sleeping during frame
}

// NewFrameRater returns a frame rate controller object for the specified
// number of target frames per second
func NewFrameRater(targetFPS uint) *FrameRater {

	f := new(FrameRater)
	f.targetDuration = time.Second / time.Duration(targetFPS)
	f.frameTimes = 0
	f.frameCount = 0
	f.lastUpdate = time.Now()
	f.timer = time.NewTimer(0)
	<-f.timer.C
	return f
}

// Start should be called at the start of the frame
func (f *FrameRater) Start() {

	f.frameStart = time.Now()
}

// Wait should be called at the end of the frame
// If necessary it will sleep to achieve the desired frame rate
func (f *FrameRater) Wait() {

	// Calculates the time duration of this frame
	elapsed := time.Now().Sub(f.frameStart)
	// Accumulates this frame time for potential FPS calculation
	f.frameCount++
	f.frameTimes += elapsed
	// If this frame duration is less than the target duration, sleeps
	// during the difference
	diff := f.targetDuration - elapsed
	if diff > 0 {
		f.timer.Reset(diff)
		<-f.timer.C
	}
}

// FPS calculates and returns the current measured FPS and the maximum
// potential FPS after the specified time interval has elapsed.
// It returns an indication if the results are valid
func (f *FrameRater) FPS(t time.Duration) (float64, float64, bool) {

	// If the time from the last update has not passed, nothing to do
	elapsed := time.Now().Sub(f.lastUpdate)
	if elapsed < t {
		return 0, 0, false
	}

	// Calculates the measured average frame rate
	fps := float64(f.frameCount) / elapsed.Seconds()
	// Calculates the average duration of a frame and the potential FPS
	frameDur := f.frameTimes.Seconds() / float64(f.frameCount)
	pfps := 1.0 / frameDur
	// Resets the frame counter and times
	f.frameCount = 0
	f.frameTimes = 0
	f.lastUpdate = time.Now()
	return fps, pfps, true
}

// UpdateFPS updates the fps value in the window title or header label
func updateFPS(ctx *Context) {

	if *oHideFPS {
		return
	}

	// Get the FPS and potential FPS from the frameRater
	fps, pfps, ok := ctx.frameRater.FPS(time.Duration(*oUpdateFPS) * time.Millisecond)
	if !ok {
		return
	}

	// Shows the values in the window title or header label
	msg := fmt.Sprintf("%3.1f / %3.1f", fps, pfps)
	if *oNogui {
		ctx.Win.SetTitle(msg)
	} else {
		ctx.labelFPS.SetText(msg)
	}
}

// winResizeEvent is called when the window resize event is received
func winResizeEvent(ctx *Context) {

	// Sets view port
	width, height := ctx.Win.GetSize()
	ctx.GS.Viewport(0, 0, int32(width), int32(height))
	aspect := float32(width) / float32(height)

	// Sets camera aspect ratio
	ctx.CamPersp.SetAspect(aspect)

	// Sets GUI root panel size
	ctx.root.SetSize(float32(width), float32(height))
}

// setupScene resets the current scene for executing a new (or first) test
func setupScene(ctx *Context) {

	// Cancel next events and clear all window subscriptions
	ctx.Win.CancelDispatch()
	ctx.Win.ClearSubscriptions()

	// Dispose of all test scene children
	ctx.Scene.DisposeChildren(true)
	if ctx.Gui != nil {
		ctx.Gui.DisposeChildren(true)
	}
	//log.Info("STATS:%+v", ctx.GS.Stats())

	// Sets default background color
	ctx.GS.ClearColor(0.6, 0.6, 0.6, 1.0)

	// Adds ambient light to the test scene
	ctx.AmbLight = light.NewAmbient(&math32.Color{1.0, 1.0, 1.0}, 0.5)
	ctx.Scene.Add(ctx.AmbLight)

	// Sets perspective camera position
	width, height := ctx.Win.GetSize()
	aspect := float32(width) / float32(height)
	if ctx.CamPersp == nil {
		ctx.CamPersp = camera.NewPerspective(65, aspect, 0.01, 1000)
	}
	ctx.CamPersp.SetPosition(0, 0, 5)
	ctx.CamPersp.LookAt(&math32.Vector3{0, 0, 0})
	ctx.CamPersp.SetAspect(aspect)

	// Sets orthographic camera
	if ctx.CamOrtho == nil {
		ctx.CamOrtho = camera.NewOrthographic(-2, 2, 2, -2, 0.01, 100)
	}
	ctx.CamOrtho.SetPosition(0, 0, 3)
	ctx.CamOrtho.LookAt(&math32.Vector3{0, 0, 0})
	ctx.CamOrtho.SetZoom(1.0)

	// Default camera is perspective
	ctx.Camera = ctx.CamPersp
	// Adds camera to the scene
	ctx.Scene.Add(ctx.Camera.GetCamera())

	// Subscribe to window key events
	ctx.Win.Subscribe(window.OnKeyDown, func(evname string, ev interface{}) {
		kev := ev.(*window.KeyEvent)
		// ESC terminates program
		if kev.Keycode == window.KeyEscape {
			ctx.Win.SetShouldClose(true)
			return
		}
		// F10 toggles full screen
		if kev.Keycode == window.KeyF10 {
			ctx.Win.SetFullScreen(!ctx.Win.FullScreen())
			return
		}
		// Ctr-Alt-S prints statistics in the console
		if kev.Keycode == window.KeyS && kev.Mods == window.ModControl|window.ModAlt {
			logStats(ctx)
		}
	})

	// Subscribe to window resize events
	ctx.Win.Subscribe(window.OnWindowSize, func(evname string, ev interface{}) {
		winResizeEvent(ctx)
	})

	// Root is the base panel for GUI
	ctx.root.SubscribeWin()

	// Creates orbit camera control
	// It is important to do this after the root panel subscription
	// to avoid GUI events being propagated to the orbit control.
	ctx.Orbit = control.NewOrbitControl(ctx.CamPersp, ctx.Win)

	// If audio active, resets global listener parameters
	if ctx.Audio {
		al.Listener3f(al.Position, 0, 0, 0)
		al.Listener3f(al.Velocity, 0, 0, 0)
		al.Listenerfv(al.Orientation, []float32{0, 0, -1, 0, 1, 0})
		// If audio capture device was opened, close it
		if ctx.CapDev != nil {
			al.CaptureStop(ctx.CapDev)
			al.CaptureCloseDevice(ctx.CapDev)
			ctx.CapDev = nil
		}
	}

	// If no gui, nothing more to do
	if ctx.Control == nil {
		return
	}

	// Remove all controls and adds default ones
	ctx.Control.Clear()

	// Adds camera selection
	cb := ctx.Control.AddCheckBox("Perspective camera")
	cb.SetValue(true)
	cb.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		if cb.Value() {
			ctx.Camera = ctx.CamPersp
		} else {
			ctx.Camera = ctx.CamOrtho
		}
		// Recreates orbit camera control
		ctx.Orbit.Dispose()
		ctx.Orbit = control.NewOrbitControl(ctx.Camera, ctx.Win)
	})

	// Adds ambient light slider
	s1 := ctx.Control.AddSlider("Ambient light:", 2.0, ctx.AmbLight.Intensity())
	s1.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		ctx.AmbLight.SetIntensity(s1.Value())
	})
}

// logStats generate log with current statistics
func logStats(ctx *Context) {

	const statsFormat = `
         Shaders: %d
            Vaos: %d
         Buffers: %d
        Textures: %d
  Uniforms/frame: %d
Draw calls/frame: %d
 CGO calls/frame: %d
`
	log.Info(statsFormat,
		ctx.stats.Glstats.Shaders,
		ctx.stats.Glstats.Vaos,
		ctx.stats.Glstats.Buffers,
		ctx.stats.Glstats.Textures,
		ctx.stats.Unisets,
		ctx.stats.Drawcalls,
		ctx.stats.Cgocalls,
	)
}

// loadAudioLibs try to load audio libraries
func loadAudioLibs(ctx *Context) {

	// Try to load OpenAL
	err := al.Load()
	if err != nil {
		log.Warn("%s", err)
		return
	}

	// Opens default audio device
	ctx.AudioDev, err = al.OpenDevice("")
	if ctx.AudioDev == nil {
		log.Warn("Error: %s opening OpenAL default device", err)
		return
	}

	// Checks for OpenAL effects extension support
	if al.IsExtensionPresent("ALC_EXT_EFX") {
		ctx.AudioEFX = true
	}

	// Creates audio context with auxiliary sends
	var attribs []int
	if ctx.AudioEFX {
		attribs = []int{al.MAX_AUXILIARY_SENDS, 4}
	}
	acx, err := al.CreateContext(ctx.AudioDev, attribs)
	if err != nil {
		log.Error("Error creating audio context:%s", err)
		return
	}

	// Makes the context the current one
	err = al.MakeContextCurrent(acx)
	if err != nil {
		log.Error("Error setting audio context current:%s", err)
		return
	}
	log.Info("%s version: %s", al.GetString(al.Vendor), al.GetString(al.Version))
	ctx.Audio = true
	if ctx.AudioEFX {
		log.Info("OpenAL EFX extension available")
	}

	// Ogg Vorbis support
	err = ov.Load()
	if err == nil {
		ctx.Vorbis = true
		vorbis.Load()
		log.Info("%s", vorbis.VersionString())
	} else {
		log.Warn("%s", err)
	}
}

// checkDirData try to find and return the complete data directory path.
// Aborts if not found
func checkDirData() string {

	// Checks first if data directory is in the current directory
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
	path := filepath.Join(execDir, "data")
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

	// Shows error message and aborts
	log.Fatal("Data directory NOT FOUND")
	return ""
}

// usage shows the application usage
func usage() {

	fmt.Fprintf(os.Stderr, "%s v%d.%d\n", ProgName, Vmajor, Vminor)
	fmt.Fprintf(os.Stderr, "usage: %s [options] [<test>] \n", ExecName)
	flag.PrintDefaults()
	os.Exit(2)
}

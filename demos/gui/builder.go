package gui

import (
	"fmt"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/util"
	"time"

	"path/filepath"
)

func init() {
	app.DemoMap["gui.builder"] = &GuiBuilder{}
}

type GuiBuilder struct {
	selFile   *util.FileSelectButton
	container *gui.Panel
}

// Start is called once at the start of the demo.
func (t *GuiBuilder) Start(a *app.App) {

	// Show and enable demo panel
	a.DemoPanel().SetRenderable(true)
	a.DemoPanel().SetEnabled(true)

	// Creates file selection button
	t.selFile = util.NewFileSelectButton(a.DirData()+"/gui", "Select File", 400, 300)
	t.selFile.SetPosition(0, 0)
	t.selFile.FS.SetFileFilters("*.yaml")
	a.DemoPanel().Add(t.selFile)
	t.selFile.Subscribe("OnSelect", func(evname string, ev interface{}) {
		fmt.Println("OnSelect")
		fpath := ev.(string)
		t.build(a, fpath)
	})
	t.selFile.SetMargins(2, 2, 2, 2)

	// Creates container
	t.container = gui.NewPanel(0, 0)
	t.container.SetBorders(0, 0, 0, 0)
	t.container.SetMargins(2, 2, 2, 2)
	t.container.SetColor4(&math32.Color4{1, 1, 1, 0})
	// Internal function to resize container when gui resizes
	onResize := func() {
		t.container.SetSize(a.DemoPanel().ContentWidth(), a.DemoPanel().ContentHeight()-t.selFile.Height())
		t.container.SetPosition(0, t.selFile.Position().Y+t.selFile.Height())
	}
	a.DemoPanel().SubscribeID(gui.OnResize, a, func(evname string, ev interface{}) { onResize() })
	a.DemoPanel().Add(t.container)
	onResize()

	// Loads default gui builder file
	t.build(a, a.DirData()+"/gui/1panels.yaml")
}

// Update is called every frame.
func (t *GuiBuilder) Update(a *app.App, deltaTime time.Duration) {}

// Cleanup is called once at the end of the demo.
func (t *GuiBuilder) Cleanup(a *app.App) {}

func (t *GuiBuilder) build(a *app.App, fpath string) {

	// Creates gui builder
	b := gui.NewBuilder()
	b.SetImagepath(a.DirData() + "/images/")

	// Parses description file
	err := b.ParseFile(fpath)
	if err != nil {
		t.selFile.Label.SetText("Select File")
		t.selFile.SetError(err.Error())
		return
	}

	t.selFile.Label.SetText("File: " + filepath.Base(fpath))
	t.selFile.SetError("")
	t.container.DisposeChildren(true)

	// Build gui objects and adds them to the container panel
	names := b.Names()
	for _, name := range names {
		p, err := b.Build(name)
		if err != nil {
			t.selFile.Label.SetText("Select File")
			t.selFile.SetError(err.Error())
			return
		}
		t.container.Add(p)
	}
}

package gui

import (
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/demos"
	"github.com/g3n/g3nd/util"

	"path/filepath"
)

func init() {
	demos.Map["gui.builder"] = &GuiBuilder{}
}

type GuiBuilder struct {
	selFile   *util.FileSelectButton
	container *gui.Panel
}

func (t *GuiBuilder) Initialize(a *app.App) {

	// Creates file selection button
	t.selFile = util.NewFileSelectButton(a.DirData()+"/gui", "Select File", 400, 300)
	t.selFile.SetPosition(0, 0)
	t.selFile.FS.SetFileFilters("*.yaml")
	a.GuiPanel().Add(t.selFile)
	t.selFile.Subscribe("OnSelect", func(evname string, ev interface{}) {
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
		t.container.SetSize(a.GuiPanel().ContentWidth(), a.GuiPanel().ContentHeight()-t.selFile.Height())
		t.container.SetPosition(0, t.selFile.Position().Y+t.selFile.Height())
	}
	a.GuiPanel().Subscribe(gui.OnResize, func(evname string, ev interface{}) { onResize() })
	a.GuiPanel().Add(t.container)
	onResize()

	// Loads default gui builder file
	t.build(a, a.DirData()+"/gui/1panels.yaml")
}

func (t *GuiBuilder) Render(app *app.App) {

}

func (t *GuiBuilder) build(app *app.App, fpath string) {

	// Creates gui builder
	b := gui.NewBuilder()
	b.SetImagepath(app.DirData() + "/images/")

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

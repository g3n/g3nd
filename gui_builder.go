package main

import (
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"path/filepath"
)

func init() {
	TestMap["gui.builder"] = &GuiBuilder{}
}

type GuiBuilder struct {
	selFile   *FileSelectButton
	container *gui.Panel
}

func (t *GuiBuilder) Initialize(ctx *Context) {

	// Creates file selection button
	t.selFile = NewFileSelectButton(ctx.DirData+"/gui", "Select File", 400, 300)
	t.selFile.SetPosition(0, 0)
	t.selFile.FS.SetFileFilters("*.yaml")
	ctx.Gui.Add(t.selFile)
	t.selFile.Subscribe("OnSelect", func(evname string, ev interface{}) {
		fpath := ev.(string)
		t.build(ctx, fpath)
	})
	t.selFile.SetMargins(2, 2, 2, 2)

	// Creates container
	t.container = gui.NewPanel(0, 0)
	t.container.SetBorders(0, 0, 0, 0)
	t.container.SetMargins(2, 2, 2, 2)
	t.container.SetColor4(&math32.Color4{1, 1, 1, 0})
	onResize := func() {
		t.container.SetSize(ctx.Gui.ContentWidth(), ctx.Gui.ContentHeight()-t.selFile.Height())
		t.container.SetPosition(0, t.selFile.Position().Y+t.selFile.Height())
	}
	ctx.Gui.Subscribe(gui.OnResize, func(evname string, ev interface{}) { onResize() })
	ctx.Gui.Add(t.container)
	onResize()

	// Loads default gui builder file
	t.build(ctx, ctx.DirData+"/gui/1panels.yaml")
}

func (t *GuiBuilder) Render(ctx *Context) {

}

func (t *GuiBuilder) build(ctx *Context, fpath string) {

	// Creates gui builder
	b := gui.NewBuilder()
	b.SetImagepath(ctx.DirData + "/images/")

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
	log.Error("NAMES:%v\n", names)
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

package main

import (
	"github.com/g3n/engine/gui"
	"path/filepath"
)

func init() {
	TestMap["gui.builder"] = &GuiBuilder{}
}

type GuiBuilder struct {
	selFile *FileSelectButton
}

func (t *GuiBuilder) Initialize(ctx *Context) {

	b := gui.NewBuilder()

	// Creates file selection button
	t.selFile = NewFileSelectButton(ctx.DirData+"/gui", "Select File", 400, 300)
	t.selFile.SetPosition(0, 0)
	t.selFile.FS.SetFileFilters("*.yaml")
	ctx.Gui.Add(t.selFile)
	t.selFile.Subscribe("OnSelect", func(evname string, ev interface{}) {
		fpath := ev.(string)
		err := b.BuildFromFile(fpath)
		if err == nil {
			t.selFile.Label.SetText("File: " + filepath.Base(fpath))
			t.selFile.SetError("")
			panels := b.Panels()
			posx := float32(0)
			posy := float32(48)
			for _, p := range panels {
				ctx.Gui.Add(p)
				p.GetPanel().SetPosition(posx, posy)
				posx += p.GetPanel().Width() + 2
			}

		} else {
			t.selFile.Label.SetText("Select File")
			t.selFile.SetError(err.Error())
		}
	})

}

func (t *GuiBuilder) Render(ctx *Context) {

}

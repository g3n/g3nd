package main

import (
	"github.com/g3n/engine/gui"
)

func init() {
	TestMap["gui.folder"] = &GuiFolder{}
}

type GuiFolder struct {
	f1 *gui.Folder
}

func (t *GuiFolder) Initialize(ctx *Context) {

	cbExpand := gui.NewCheckBox("Expand left")
	cbExpand.SetPosition(200, 10)
	cbExpand.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		t.f1.SetAlignRight(!cbExpand.Value())
	})
	ctx.Gui.Add(cbExpand)

	tree := gui.NewTree(260, 300)
	tree.Add(gui.NewImageLabel("A folder shows/hides an enclosed panel"))
	tree.Add(gui.NewCheckBox("checkbox 1"))
	tree.Add(gui.NewRadioButton("radio button 1"))
	n1 := tree.AddNode("Node 1")
	n1.Add(gui.NewImageLabel("image label 2"))
	n1.Add(gui.NewCheckBox("checkbox 2"))
	n1.Add(gui.NewRadioButton("radio button 2"))
	n2 := n1.AddNode("Node 2")
	n2.Add(gui.NewImageLabel("image label 3"))
	n2.Add(gui.NewCheckBox("checkbox 3"))
	n2.Add(gui.NewRadioButton("radio button 3"))

	t.f1 = gui.NewFolder("folder1", 200, tree)
	t.f1.SetPosition(200, 40)
	t.f1.SetAlignRight(true)
	ctx.Gui.Add(t.f1)
}

func (t *GuiFolder) Render(ctx *Context) {
}

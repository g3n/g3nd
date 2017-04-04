package main

func init() {
	TestMap["gui.menu"] = &GuiMenu{}
}

type GuiMenu struct {
}

func (t *GuiMenu) Initialize(ctx *Context) {
}

func (t *GuiMenu) Render(ctx *Context) {
}

package main

import (
	"github.com/g3n/engine/gui"
)

func init() {
	TestMap["gui.table"] = &GuiTable{}
}

type GuiTable struct{}

func (t *GuiTable) Initialize(ctx *Context) {

	// Creates Table
	tableY := float32(20)
	tab, err := gui.NewTable(400, 200, []gui.TableColumn{
		{Id: "1", Name: "Col1", Width: 100, Format: "US$%v"},
		{Id: "2", Name: "Col2", Width: 80},
		{Id: "3", Name: "Col3", Width: 120},
	})
	if err != nil {
		panic(err)
	}

	// Sets the table data
	tab.SetRows([]map[string]interface{}{
		{"1": 10, "2": 20, "3": "Text"},
		{"1": 20, "2": 30, "3": "long text message"},
	})
	ctx.Gui.Add(tab)

	tab.SetBorders(1, 1, 1, 1)
	tab.SetPosition(0, tableY)
	tab.SetMargins(10, 10, 10, 10)
	tab.SetSize(ctx.Gui.ContentWidth(), ctx.Gui.ContentHeight()-tableY)
	ctx.Gui.Subscribe(gui.OnResize, func(evname string, ev interface{}) {
		tab.SetSize(ctx.Gui.ContentWidth(), ctx.Gui.ContentHeight()-tableY)
	})

}

func (t *GuiTable) Render(ctx *Context) {

}

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
	tab, err := gui.NewTable(400, 200, []gui.TableColumn{
		{Id: "1", Name: "Col1", Width: 100, Format: "US$%v"},
		{Id: "2", Name: "Col2", Width: 80},
		{Id: "3", Name: "Col3", Width: 0},
	})
	if err != nil {
		panic(err)
	}
	tab.SetBorders(1, 1, 1, 1)
	tab.SetPosition(10, 10)

	// Sets the table data
	tab.SetRows([]map[string]interface{}{
		{"1": 10, "2": 20},
		{"1": 20, "2": 30},
	})
	ctx.Gui.Add(tab)

}

func (t *GuiTable) Render(ctx *Context) {

}

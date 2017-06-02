package main

import (
	"strconv"
	"time"

	"github.com/g3n/engine/gui"
)

func init() {
	TestMap["gui.table"] = &GuiTable{}
}

type GuiTable struct {
}

func (t *GuiTable) Initialize(ctx *Context) {

	// Generates data rows
	nextRow := 0
	genRows := func(count int) []map[string]interface{} {

		values := make([]map[string]interface{}, 0, count)
		for i := 0; i < count; i++ {
			rval := make(map[string]interface{})
			rval["1"] = strconv.Itoa(nextRow)
			rval["2"] = float64(nextRow) / 3
			rval["3"] = time.Now()
			values = append(values, rval)
			nextRow++
		}
		return values
	}

	// Create menu with test options
	mb := gui.NewMenuBar()
	mb.SetPosition(10, 10)

	// Create Header Menu
	mh := gui.NewMenu()
	mh.AddOption("Show header").SetId("showHeader")
	mh.AddOption("Hide header").SetId("hideHeader")
	mb.AddMenu("Header", mh)

	// Create Row Menu
	mr := gui.NewMenu()
	mr.AddOption("Add row").SetId("addRow")
	mr.AddOption("Add 10 rows").SetId("add10Rows")
	mr.AddSeparator()
	mr.AddOption("Insert row").SetId("insRow")
	mr.AddOption("Insert 10 rows").SetId("ins10Rows")
	mb.AddMenu("Row", mr)
	mr.AddSeparator()
	mr.AddOption("Remove top row").SetId("remTopRow")
	mr.AddOption("Remove 10 top rows").SetId("rem10TopRows")
	mr.AddOption("Remove bottom row").SetId("remBottomRow")
	mr.AddOption("Remove 10 bottom rows").SetId("rem10BottomRows")

	ctx.Gui.Add(mb)
	ctx.Gui.Root().SetKeyFocus(mb)

	// Creates Table
	tableY := mb.Position().Y + mb.Height() + 10
	tab, err := gui.NewTable(400, 200, []gui.TableColumn{
		{Id: "1", Name: "Col1", Width: 100},
		{Id: "2", Name: "Col2", Width: 100},
		{Id: "3", Name: "Col3", Width: 100},
	})
	if err != nil {
		panic(err)
	}

	// Sets the table data
	tab.SetRows(genRows(2))
	ctx.Gui.Add(tab)

	tab.SetBorders(1, 1, 1, 1)
	tab.SetPosition(0, tableY)
	tab.SetMargins(10, 10, 10, 10)
	tab.SetSize(ctx.Gui.ContentWidth(), ctx.Gui.ContentHeight()-tableY)
	ctx.Gui.Subscribe(gui.OnResize, func(evname string, ev interface{}) {
		tab.SetSize(ctx.Gui.ContentWidth(), ctx.Gui.ContentHeight()-tableY)
	})

	mb.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		opid := ev.(*gui.MenuItem).Id()
		switch opid {
		case "showHeader":
			tab.ShowHeader(true)
		case "hideHeader":
			tab.ShowHeader(false)
		case "addRow":
			tab.AddRow(genRows(1)[0])
		case "add10Rows":
			values := genRows(10)
			for i := 0; i < len(values); i++ {
				tab.AddRow(values[i])
			}
		case "insRow":
			tab.InsertRow(0, genRows(1)[0])
		case "ins10Rows":
			values := genRows(10)
			for i := 0; i < len(values); i++ {
				tab.InsertRow(0, values[i])
			}
		case "remTopRow":
			if tab.Len() > 0 {
				tab.RemoveRow(0)
			}
		case "rem10TopRows":
			count := 10
			for count > 0 && tab.Len() > 0 {
				tab.RemoveRow(0)
				count--
			}
		case "remBottomRow":
			if tab.Len() > 0 {
				tab.RemoveRow(tab.Len() - 1)
			}
		case "rem10BottomRows":
			count := 10
			for count > 0 && tab.Len() > 0 {
				tab.RemoveRow(tab.Len() - 1)
				count--
			}
		}
	})

}

func (t *GuiTable) Render(ctx *Context) {

}

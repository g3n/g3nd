package main

import (
	"fmt"
	"time"

	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/window"
)

func init() {
	TestMap["gui.table"] = &GuiTable{}
}

type GuiTable struct {
}

func (t *GuiTable) Initialize(ctx *Context) {

	// Function to generate table data rows
	nextRow := 0
	genRows := func(count int) []map[string]interface{} {

		values := make([]map[string]interface{}, 0, count)
		for i := 0; i < count; i++ {
			rval := make(map[string]interface{})
			rval["1"] = nextRow // column id
			rval["2"] = time.Now()
			rval["3"] = float64(nextRow) / 3
			rval["4"] = ""
			values = append(values, rval)
			nextRow++
		}
		return values
	}

	// Create menu with test options
	mb := gui.NewMenuBar()
	mb.SetPosition(10, 10)

	// Create Header Menu
	mt := gui.NewMenu()
	mt.AddOption("Show header").SetId("showHeader")
	mt.AddOption("Hide header").SetId("hideHeader")
	mt.AddSeparator()
	mt.AddOption("Show all columns").SetId("showAllColumns")
	mt.AddSeparator()
	mt.AddOption("Show status").SetId("showStatus")
	mt.AddOption("Hide status").SetId("hideStatus")
	mb.AddMenu("Table", mt)

	// Create Row Menu
	mr := gui.NewMenu()
	mr.AddOption("Add row").SetId("addRow")
	mr.AddOption("Add 10 rows").SetId("add10Rows")
	mr.AddOption("Add 50 rows").SetId("add50Rows")
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

	// Time formatting function
	//formatTime := func(row int, col string, val interface{}) string {
	formatTime := func(cell gui.TableCell) string {
		if cell.Value == nil {
			return ""
		}
		t := cell.Value.(time.Time)
		return time.Time(t).Format("15:04:05.000")
	}

	// Formatting function for calculated column 4
	formatCalc := func(cell gui.TableCell) string {
		c1 := cell.Tab.Cell("1", cell.Row)
		if c1 == nil {
			return ""
		}
		v := c1.(int)
		return fmt.Sprintf("Col1 / 5 = %6.2f", float64(v)/5)
	}

	// Creates Table
	tableY := mb.Position().Y + mb.Height() + 10
	tab, err := gui.NewTable(400, 200, []gui.TableColumn{
		{Id: "1", Name: "Col1", Width: 40, Align: gui.AlignLeft, Format: "%04d"},
		{Id: "2", Name: "Col2", Width: 100, Align: gui.AlignCenter, FormatFunc: formatTime},
		{Id: "3", Name: "Col3", Width: 100, Align: gui.AlignRight, Format: "US$%6.2f"},
		{Id: "4", Name: "Col4", Width: 140, Hidden: true, Align: gui.AlignCenter, FormatFunc: formatCalc},
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

	// Creates column context menu
	mCol := gui.NewMenu()
	mCol.AddOption("Hide column").SetId("hideColumn")
	mCol.AddOption("Show all columns").SetId("showAllColumns")
	mCol.AddSeparator()
	mCol.AddOption("Move column left").SetId("moveColumnLeft")
	mCol.SetVisible(false)
	mCol.SetBounded(false)
	tab.Add(mCol)

	// Creates row context menu
	mRow := gui.NewMenu()
	mRow.AddOption("Delete row").SetId("delRow")
	mRow.AddSeparator()
	mRow.AddOption("Insert row above").SetId("insRowAbove")
	mRow.AddOption("Insert row below").SetId("insRowBelow")
	mRow.SetVisible(false)
	mRow.SetBounded(false)
	tab.Add(mRow)

	// Subscribe to table on click
	var tce gui.TableClickEvent
	tab.Subscribe(gui.OnTableClick, func(evname string, ev interface{}) {
		tce = ev.(gui.TableClickEvent)
		mRow.SetVisible(false)
		mCol.SetVisible(false)
		//log.Debug("evname:%v -> %+v", evname, e)
		if tce.Button != window.MouseButtonRight {
			return
		}
		if tce.Header {
			mCol.SetPosition(tce.X, tce.Y)
			mCol.SetVisible(true)
			return
		}
		if tce.Row >= 0 {
			mRow.SetPosition(tce.X, tce.Y)
			mRow.SetVisible(true)
			return
		}
	})

	// Subscribe to table row count event
	tab.Subscribe(gui.OnTableRowCount, func(evname string, ev interface{}) {
		selRow := tab.SelectedRow()
		tab.SetStatusText(fmt.Sprintf("Count:%d Selected:%d", tab.RowCount(), selRow))
	})

	// Subscribe to table onchange event
	tab.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		selRow := tab.SelectedRow()
		tab.SetStatusText(fmt.Sprintf("Count:%d Selected:%d", tab.RowCount(), selRow))
	})

	mCol.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		mCol.SetVisible(false)
		opid := ev.(*gui.MenuItem).Id()
		switch opid {
		case "hideColumn":
			tab.ShowColumn(tce.Col, false)
		case "showAllColumns":
			tab.ShowAllColumns()
		}
	})

	mRow.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		mRow.SetVisible(false)
		opid := ev.(*gui.MenuItem).Id()
		switch opid {
		case "delRow":
			tab.RemoveRow(tce.Row)
		case "insRowAbove":
			values := genRows(1)
			tab.InsertRow(tce.Row, values[0])
		case "insRowBelow":
			values := genRows(1)
			tab.InsertRow(tce.Row+1, values[0])
		}
	})

	mb.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		opid := ev.(*gui.MenuItem).Id()
		switch opid {
		case "showHeader":
			tab.ShowHeader(true)
		case "hideHeader":
			tab.ShowHeader(false)
		case "showAllColumns":
			tab.ShowAllColumns()
		case "showStatus":
			tab.ShowStatus(true)
		case "hideStatus":
			tab.ShowStatus(false)
		case "addRow":
			tab.AddRow(genRows(1)[0])
		case "add10Rows":
			values := genRows(10)
			for i := 0; i < len(values); i++ {
				tab.AddRow(values[i])
			}
		case "add50Rows":
			values := genRows(50)
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
			if tab.RowCount() > 0 {
				tab.RemoveRow(0)
			}
		case "rem10TopRows":
			count := 10
			for count > 0 && tab.RowCount() > 0 {
				tab.RemoveRow(0)
				count--
			}
		case "remBottomRow":
			if tab.RowCount() > 0 {
				tab.RemoveRow(tab.RowCount() - 1)
			}
		case "rem10BottomRows":
			count := 10
			for count > 0 && tab.RowCount() > 0 {
				tab.RemoveRow(tab.RowCount() - 1)
				count--
			}
		}
	})

}

func (t *GuiTable) Render(ctx *Context) {

}

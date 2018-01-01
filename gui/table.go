package gui

import (
	"fmt"
	"strconv"
	"time"

	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/window"
	"github.com/g3n/g3nd/demos"
	"github.com/g3n/g3nd/g3nd"
)

func init() {
	demos.Map["gui.table"] = &GuiTable{}
}

type GuiTable struct {
}

func (t *GuiTable) Initialize(app *g3nd.App) {

	// Function to generate table data rows
	nextRow := 0
	genRows := func(count int) []map[string]interface{} {

		values := make([]map[string]interface{}, 0, count)
		for i := 0; i < count; i++ {
			rval := make(map[string]interface{})
			rval["1"] = nextRow // column id
			rval["2"] = time.Now()
			rval["3"] = float64(nextRow) / 3
			rval["4"] = nil
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
	mt.AddSeparator()
	mt.AddOption("Enable multi row selection").SetId("enableMultirow")
	mt.AddOption("Disable multi row selection").SetId("disableMultirow")
	mb.AddMenu("Table", mt)

	// Create Row Menu
	mr := gui.NewMenu()
	mr.AddOption("Add row").SetId("addRow")
	mr.AddOption("Add 10 rows").SetId("add10Rows")
	mr.AddOption("Add 100 rows").SetId("add100Rows")
	mr.AddSeparator()
	mr.AddOption("Remove top row").SetId("remTopRow")
	mr.AddOption("Remove 100 top rows").SetId("rem100TopRows")
	mr.AddSeparator()
	mr.AddOption("Remove bottom row").SetId("remBottomRow")
	mr.AddOption("Remove 100 bottom rows").SetId("rem100BottomRows")
	mr.AddSeparator()
	mr.AddOption("Remove all rows").SetId("remAllRows")
	mb.AddMenu("Row", mr)

	app.GuiPanel().Add(mb)

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
		{Id: "1", Header: "Col1", Width: 100, Minwidth: 32, Align: gui.AlignLeft, Format: "%d", Resize: true, Sort: gui.TableSortNumber, Expand: 0},
		{Id: "2", Header: "Col2", Width: 100, Minwidth: 32, Align: gui.AlignCenter, FormatFunc: formatTime, Resize: true, Expand: 0},
		{Id: "3", Header: "Col3", Width: 100, Minwidth: 32, Align: gui.AlignRight, Format: "US$%6.2f", Resize: true, Expand: 0},
		{Id: "4", Header: "Col4", Width: 100, Minwidth: 32, Align: gui.AlignCenter, FormatFunc: formatCalc, Expand: 0},
	})
	if err != nil {
		panic(err)
	}
	// Sets the table data
	tab.SetRows(genRows(2))
	tab.SetBorders(1, 1, 1, 1)
	tab.SetPosition(0, tableY)
	tab.SetMargins(10, 10, 10, 10)
	tab.SetSize(app.GuiPanel().ContentWidth(), app.GuiPanel().ContentHeight()-tableY)
	app.GuiPanel().Add(tab)

	// Resize table
	app.GuiPanel().Subscribe(gui.OnResize, func(evname string, ev interface{}) {
		tab.SetSize(app.GuiPanel().ContentWidth(), app.GuiPanel().ContentHeight()-tableY)
	})

	// Creates column context menu
	mCol := gui.NewMenu()
	mCol.AddOption("Hide column").SetId("hideColumn")
	mCol.AddOption("Show all columns").SetId("showAllColumns")
	mCol.AddSeparator()
	mCol.AddOption("Move column left").SetId("moveColumnLeft")
	mCol.AddOption("Move column right").SetId("moveColumnRight")
	mCol.AddSeparator()
	mCol.AddOption("Enable resize").SetId("enableColResize")
	mCol.AddOption("Disable resize").SetId("disableColResize")
	mCol.AddSeparator()
	mCol.AddOption("Set expand to 0").SetId("setExpand0")
	mCol.AddOption("Set expand to 1").SetId("setExpand1")
	mCol.AddOption("Set expand to 2").SetId("setExpand2")
	mCol.SetVisible(false)
	mCol.SetBounded(false)
	mCol.Subscribe(gui.OnMouseOut, func(evname string, ev interface{}) {
		mCol.SetVisible(false)
	})
	tab.Add(mCol)

	// Creates row context menu
	mRow := gui.NewMenu()
	mRow.AddOption("Delete row").SetId("delRow")
	mRow.AddSeparator()
	mRow.AddOption("Insert row above").SetId("insRowAbove")
	mRow.AddOption("Insert row below").SetId("insRowBelow")
	mRow.SetVisible(false)
	mRow.SetBounded(false)
	mRow.Subscribe(gui.OnMouseOut, func(evname string, ev interface{}) {
		mRow.SetVisible(false)
	})
	tab.Add(mRow)

	// Subscribe to table on click
	var tce gui.TableClickEvent
	tab.Subscribe(gui.OnTableClick, func(evname string, ev interface{}) {
		tce = ev.(gui.TableClickEvent)
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

	// Subscribe to events to update the table status line
	updateStatus := func(evname string, ev interface{}) {
		selRows := tab.SelectedRows()
		current := ""
		if len(selRows) > 0 {
			current = strconv.Itoa(selRows[0])
		}
		tab.SetStatusText(fmt.Sprintf("Rows: %d Selected: %d Current: %s", tab.RowCount(), len(selRows), current))
	}
	tab.Subscribe(gui.OnTableRowCount, updateStatus)
	tab.Subscribe(gui.OnChange, updateStatus)

	// Subscribe to column menu
	mCol.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		mCol.SetVisible(false)
		opid := ev.(*gui.MenuItem).Id()
		switch opid {
		case "hideColumn":
			tab.ShowColumn(tce.Col, false)
		case "showAllColumns":
			tab.ShowAllColumns()
		case "moveColumnLeft":
			if tce.ColOrder >= 1 {
				tab.SetColOrder(tce.Col, tce.ColOrder-1)
			}
		case "moveColumnRight":
			if tce.ColOrder < 3 {
				tab.SetColOrder(tce.Col, tce.ColOrder+1)
			}
		case "enableColResize":
			tab.EnableColResize(tce.Col, true)
		case "disableColResize":
			tab.EnableColResize(tce.Col, false)
		case "setExpand0":
			tab.SetColExpand(tce.Col, 0)
		case "setExpand1":
			tab.SetColExpand(tce.Col, 1)
		case "setExpand2":
			tab.SetColExpand(tce.Col, 2)
		}
	})

	// Subscribe to row context menu
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
		case "enableMultirow":
			tab.SetSelectionType(gui.TableSelMultiRow)
		case "disableMultirow":
			tab.SetSelectionType(gui.TableSelSingleRow)
		case "addRow":
			tab.AddRow(genRows(1)[0])
		case "add10Rows":
			values := genRows(10)
			for i := 0; i < len(values); i++ {
				tab.AddRow(values[i])
			}
		case "add100Rows":
			values := genRows(100)
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
		case "ins100Rows":
			values := genRows(100)
			for i := 0; i < len(values); i++ {
				tab.InsertRow(0, values[i])
			}
		case "remTopRow":
			if tab.RowCount() > 0 {
				tab.RemoveRow(0)
			}
		case "rem100TopRows":
			count := 100
			for count > 0 && tab.RowCount() > 0 {
				tab.RemoveRow(0)
				count--
			}
		case "remBottomRow":
			if tab.RowCount() > 0 {
				tab.RemoveRow(tab.RowCount() - 1)
			}
		case "rem100BottomRows":
			count := 100
			for count > 0 && tab.RowCount() > 0 {
				tab.RemoveRow(tab.RowCount() - 1)
				count--
			}
		case "remAllRows":
			tab.Clear()
		}
	})

}

func (t *GuiTable) Render(app *g3nd.App) {

}

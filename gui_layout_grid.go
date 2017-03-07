package main

import (
	"fmt"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/graphic"
)

func init() {
	TestMap["gui.layout_grid"] = &GuiLayoutGrid{}
}

type GuiLayoutGrid struct{}

func (t *GuiLayoutGrid) Initialize(ctx *Context) {

	axis := graphic.NewAxisHelper(1)
	ctx.Scene.Add(axis)

	p1 := gui.NewPanel(300, 200)
	p1.SetPosition(10, 10)
	p1.SetColor(&math32.White)
	l1 := gui.NewGridLayout()
	p1.SetLayout(l1)
	ctx.Gui.Add(p1)
	for row := 0; row < 4; row++ {
		for col := 0; col < 5; col++ {
			text := fmt.Sprintf("item %d.%d", row, col)
			l := gui.NewLabel(text)
			l.SetPaddings(4, 4, 4, 4)
			l.SetLayoutParams(&gui.GridLayoutParams{Row: row, Col: col})
			p1.Add(l)
		}
	}

	p2 := gui.NewPanel(300, 300)
	p2.SetPosition(10, 250)
	p2.SetColor(&math32.White)
	l2 := gui.NewGridLayout()
	p2.SetLayout(l2)
	ctx.Gui.Add(p2)
	type gridItem struct {
		w       float32
		h       float32
		color   *math32.Color
		row     int
		col     int
		colSpan int
		valign  gui.Align
		halign  gui.Align
	}
	grid := []gridItem{
		{20, 20, &math32.Black, 0, 0, 0, gui.AlignNone, gui.AlignNone},
		{30, 30, &math32.Red, 0, 1, 0, gui.AlignNone, gui.AlignNone},
		{40, 40, &math32.Green, 0, 2, 0, gui.AlignNone, gui.AlignNone},
		{50, 50, &math32.Blue, 0, 3, 0, gui.AlignNone, gui.AlignNone},
		{60, 60, &math32.Gray, 0, 4, 0, gui.AlignNone, gui.AlignCenter},

		{60, 60, &math32.Gray, 1, 0, 0, gui.AlignBottom, gui.AlignRight},
		{50, 50, &math32.Blue, 1, 1, 0, gui.AlignBottom, gui.AlignRight},
		{40, 40, &math32.Green, 1, 2, 0, gui.AlignBottom, gui.AlignRight},
		{30, 30, &math32.Red, 1, 3, 0, gui.AlignBottom, gui.AlignRight},
		{20, 20, &math32.Black, 1, 4, 0, gui.AlignBottom, gui.AlignRight},

		{20, 20, &math32.Black, 2, 0, 0, gui.AlignCenter, gui.AlignCenter},
		{30, 30, &math32.Red, 2, 1, 0, gui.AlignCenter, gui.AlignCenter},
		{40, 40, &math32.Green, 2, 2, 0, gui.AlignCenter, gui.AlignCenter},
		{50, 50, &math32.Blue, 2, 3, 0, gui.AlignCenter, gui.AlignCenter},
		{60, 60, &math32.Gray, 2, 4, 0, gui.AlignCenter, gui.AlignCenter},

		{60, 60, &math32.Gray, 3, 0, 0, gui.AlignTop, gui.AlignRight},
		{50, 50, &math32.Blue, 3, 1, 0, gui.AlignTop, gui.AlignRight},
		{40, 40, &math32.Green, 3, 2, 0, gui.AlignTop, gui.AlignRight},
		{30, 30, &math32.Red, 3, 3, 0, gui.AlignTop, gui.AlignRight},
		{20, 20, &math32.Black, 3, 4, 0, gui.AlignTop, gui.AlignRight},

		{20, 20, &math32.Black, 4, 0, 0, gui.AlignCenter, gui.AlignCenter},
		{40, 20, &math32.Red, 4, 1, 3, gui.AlignCenter, gui.AlignCenter},
	}
	for _, item := range grid {
		p := gui.NewPanel(item.w, item.h)
		p.SetColor(item.color)
		p.SetLayoutParams(&gui.GridLayoutParams{Row: item.row, Col: item.col, ColSpan: item.colSpan, AlignV: item.valign, AlignH: item.halign})
		p2.Add(p)
	}
}

func (t *GuiLayoutGrid) Render(ctx *Context) {
}



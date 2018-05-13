package gui

import (
	"fmt"

	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/demos"
)

func init() {
	demos.Map["gui.itemscroller"] = &GuiScroller{}
}

type GuiScroller struct{}

func (t *GuiScroller) Initialize(a *app.App) {

	// Scroller 1
	s1 := gui.NewVScroller(100, 200)
	s1.SetPosition(10, 10)
	a.GuiPanel().Add(s1)
	// Scroller 1 - add button
	b1 := gui.NewButton("Add")
	b1.SetPosition(s1.Position().X+s1.Width()+10, s1.Position().Y)
	b1.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		if s1.Len()%3 == 0 {
			s1.Add(gui.NewButton(fmt.Sprintf("button %d", s1.Len())))
			return
		}
		if s1.Len()%2 == 0 {
			text := fmt.Sprintf("label %d", s1.Len())
			s1.Add(gui.NewLabel(text))
			return
		}
		img, err := gui.NewImage(a.DirData() + "/images/ok.png")
		if err != nil {
			a.Log().Fatal("%s", err)
		}
		s1.Add(img)
	})
	a.GuiPanel().Add(b1)
	// Scroller 1 - remove button
	b2 := gui.NewButton("Del")
	b2.SetPosition(s1.Position().X+s1.Width()+10, s1.Position().Y+30)
	b2.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		if s1.Len() > 0 {
			p := s1.RemoveAt(0)
			p.Dispose()
		}
	})
	a.GuiPanel().Add(b2)

	// Scroller 2
	s2 := gui.NewHScroller(300, 100)
	s2.SetPosition(10, 240)
	a.GuiPanel().Add(s2)
	// Scroller 2 - add button
	b21 := gui.NewButton("Add")
	b21.SetPosition(s2.Position().X, s2.Position().Y+s2.Height()+10)
	a.GuiPanel().Add(b21)
	b21.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		l := gui.NewLabel(" ")
		l.SetPaddings(2, 4, 2, 4)
		l.SetBordersColor(math32.NewColor("black"))
		l.SetText(fmt.Sprintf("label %d", s2.Len()))
		s2.Add(l)
		return
	})
	// Scroller 2 - remove button
	b22 := gui.NewButton("Del")
	b22.SetPosition(s2.Position().X+50, s2.Position().Y+s2.Height()+10)
	b22.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		if s2.Len() > 0 {
			s2.RemoveAt(0)
		}
	})
	a.GuiPanel().Add(b22)
}

func (t *GuiScroller) Render(a *app.App) {
}

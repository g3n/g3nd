package main

import (
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/gui/assets/icon"
	"github.com/g3n/engine/math32"
)

func init() {
	TestMap["gui.button"] = &GuiButton{}
}

type GuiButton struct{}

func (t *GuiButton) Initialize(ctx *Context) {

	axis := graphic.NewAxisHelper(1)
	ctx.Scene.Add(axis)

	// Button 1
	b1 := gui.NewButton("button 1")
	b1.SetPosition(10, 10)
	b1.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		log.Info("button 1 OnClick")
	})
	ctx.Gui.Add(b1)

	// Button 2
	b2 := gui.NewButton("button 2")
	b2.SetPosition(100, 10)
	b2.Label.SetFontSize(20)
	b2.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		log.Info("button 2 OnClick")
	})
	ctx.Gui.Add(b2)

	// Button 3
	b3 := gui.NewButton("button 3")
	b3.SetPosition(200, 10)
	b3.Label.SetFontSize(22)
	b3.SetIcon(icon.Check)
	b3.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		log.Info("button 3 OnClick")
	})
	ctx.Gui.Add(b3)

	// Button 4
	b4 := gui.NewButton("button 4")
	b4.SetPosition(340, 10)
	b4.Label.SetFontSize(24)
	b4.SetImage(ctx.DirData + "/images/ok.png")
	b4.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		log.Info("button 4 OnClick")
	})
	b4.SetSize(150, 100)
	ctx.Gui.Add(b4)

	// Button 5
	b5 := gui.NewButton("Button 5 (style changed)")
	b5.SetPosition(10, 200)
	b5.Label.SetFontSize(18)
	b5.SetIcon(icon.ViewHeadline)
	// Copy and change the default style
	styles := gui.StyleDefault().Button
	styles.Over.BorderColor = math32.Color4Name("red", 1)
	styles.Pressed.BorderColor = math32.Color4Name("red", 1)
	styles.Normal.FgColor = math32.ColorName("blue")
	b5.SetStyles(&styles)
	b5.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		log.Info("button 5 OnClick")
	})
	ctx.Gui.Add(b5)

	// Button 6
	b6Enabled := "Button 6 (enabled)"
	b6Disabled := "Button 6 (disabled)"
	b6 := gui.NewButton(b6Disabled)
	b6.SetEnabled(false)
	b6.SetPosition(300, 200)
	//b6.Label.SetFontSize(16)
	b6.SetIcon(icon.TrendingFlat)
	b6.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		log.Info("button 6 OnClick")
	})
	b6.Subscribe(gui.OnEnable, func(name string, ev interface{}) {
		log.Info("button 6 OnEnable")
		if b6.Enabled() {
			b6.Label.SetText(b6Enabled)
		} else {
			b6.Label.SetText(b6Disabled)
		}
	})
	ctx.Gui.Add(b6)

	// Button 7
	b7 := gui.NewButton("Button 7 (control)")
	b7.SetPosition(480, 200)
	b7.Label.SetFontSize(16)
	b7.SetIcon(icon.TrendingFlat)
	b7.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		log.Info("button 7 OnClick")
		b6.SetEnabled(!b6.Enabled())
	})
	ctx.Gui.Add(b7)
}

func (t *GuiButton) Render(ctx *Context) {

}

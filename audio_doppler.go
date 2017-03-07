package main

import (
	"github.com/g3n/engine/audio"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
)

func init() {
	TestMap["audio.doppler"] = &AudioDoppler{}
}

type AudioDoppler struct {
	ps1 *PlayerSphere
	ps2 *PlayerSphere
	ps3 *PlayerSphere
	ps4 *PlayerSphere
	ps5 *PlayerSphere
	ps6 *PlayerSphere
}

func (t *AudioDoppler) Initialize(ctx *Context) {

	// Show axis helper
	ah := graphic.NewAxisHelper(1.0)
	ctx.Scene.Add(ah)

	// Show grid helper
	grid := graphic.NewGridHelper(100, 1, &math32.Color{0.4, 0.4, 0.4})
	ctx.Scene.Add(grid)

	// Sets camera position
	ctx.Camera.GetCamera().SetPosition(0, 4, 12)

	// Creates listener and adds it to the current camera
	listener := audio.NewListener()
	ctx.Camera.GetCamera().Add(listener)

	// Creates player sphere
	t.ps1 = NewPlayerSphere(ctx, "engine.ogg", &math32.Color{1, 0, 0})
	t.ps1.SetPosition(-3, 0, 50)
	t.ps1.speed = -20.00
	t.ps1.player.SetRolloffFactor(1)
	ctx.Scene.Add(t.ps1)

	t.ps2 = NewPlayerSphere(ctx, "tone_1khz.wav", &math32.Color{0, 1, 0})
	t.ps2.SetPosition(-2, 0, 50)
	t.ps2.speed = -25.00
	ctx.Scene.Add(t.ps2)

	t.ps3 = NewPlayerSphere(ctx, "tone_2khz.wav", &math32.Color{0, 0, 1})
	t.ps3.SetPosition(-1, 0, 50)
	t.ps3.speed = -30.00
	ctx.Scene.Add(t.ps3)

	t.ps4 = NewPlayerSphere(ctx, "engine.ogg", &math32.Color{0, 1, 1})
	t.ps4.SetPosition(1, 0, -50)
	t.ps4.speed = 20.00
	ctx.Scene.Add(t.ps4)

	t.ps5 = NewPlayerSphere(ctx, "tone_1khz.wav", &math32.Color{1, 0, 1})
	t.ps5.SetPosition(2, 0, -50)
	t.ps5.speed = 25.00
	ctx.Scene.Add(t.ps5)

	t.ps6 = NewPlayerSphere(ctx, "tone_2khz.wav", &math32.Color{1, 1, 1})
	t.ps6.SetPosition(2, 0, -50)
	t.ps6.speed = 30.00
	ctx.Scene.Add(t.ps6)

	// Add controls
	if ctx.Control == nil {
		return
	}
	g := ctx.Control.AddGroup("Play sources")
	cb1 := g.AddCheckBox("engine -Z").SetValue(true)
	cb1.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		t.ps1.Toggle()
	})
	cb2 := g.AddCheckBox("tone_1khz -Z").SetValue(true)
	cb2.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		t.ps2.Toggle()
	})
	cb3 := g.AddCheckBox("tone_2khz -Z").SetValue(true)
	cb3.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		t.ps3.Toggle()
	})
	cb4 := g.AddCheckBox("engine +Z").SetValue(true)
	cb4.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		t.ps4.Toggle()
	})
	cb5 := g.AddCheckBox("tone_1khz +Z").SetValue(true)
	cb5.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		t.ps5.Toggle()
	})
	cb6 := g.AddCheckBox("tone_2khz +Z").SetValue(true)
	cb6.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		t.ps6.Toggle()
	})
}

func (t *AudioDoppler) Render(ctx *Context) {

	t.ps1.UpdateVel(ctx)
	t.ps2.UpdateVel(ctx)
	t.ps3.UpdateVel(ctx)
	t.ps4.UpdateVel(ctx)
	t.ps5.UpdateVel(ctx)
	t.ps6.UpdateVel(ctx)
}

func (ps *PlayerSphere) UpdateVel(ctx *Context) {

	delta := float32(ctx.TimeDelta.Seconds()) * ps.speed
	pos := ps.Position()
	pos.Z += delta
	if pos.Z >= 100 {
		pos.Z = -50
	}
	if pos.Z <= -100 {
		pos.Z = 50
	}
	ps.player.SetVelocity(0, 0, ps.speed)
	ps.SetPositionVec(&pos)

}

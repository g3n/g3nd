// +build !wasm

package audio

import (
	"github.com/g3n/engine/audio"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/util"
	"github.com/g3n/g3nd/app"
	"time"
)

func init() {
	app.DemoMap["audio.doppler"] = &AudioDoppler{}
}

type AudioDoppler struct {
	ps1 *PlayerSphere
	ps2 *PlayerSphere
	ps3 *PlayerSphere
	ps4 *PlayerSphere
	ps5 *PlayerSphere
	ps6 *PlayerSphere
}

// Start is called once at the start of the demo.
func (t *AudioDoppler) Start(a *app.App) {

	// Show axis helper
	ah := util.NewAxisHelper(1.0)
	a.Scene().Add(ah)

	// Show grid helper
	grid := util.NewGridHelper(100, 1, &math32.Color{0.4, 0.4, 0.4})
	a.Scene().Add(grid)

	// Sets camera position
	a.Camera().SetPosition(0, 4, 12)
	pos := a.Camera().Position()
	a.Camera().UpdateSize(pos.Length())
	a.Camera().LookAt(&math32.Vector3{0, 0, 0}, &math32.Vector3{0, 1, 0})

	// Creates listener and adds it to the current camera
	listener := audio.NewListener()
	a.Camera().Add(listener)

	// Creates player sphere
	t.ps1 = NewPlayerSphere(a, "engine.ogg", &math32.Color{1, 0, 0})
	t.ps1.SetPosition(-3, 0, 50)
	t.ps1.speed = -20.00
	t.ps1.player.SetRolloffFactor(1)
	a.Scene().Add(t.ps1)

	t.ps2 = NewPlayerSphere(a, "tone_1khz.wav", &math32.Color{0, 1, 0})
	t.ps2.SetPosition(-2, 0, 50)
	t.ps2.speed = -25.00
	a.Scene().Add(t.ps2)

	t.ps3 = NewPlayerSphere(a, "tone_2khz.wav", &math32.Color{0, 0, 1})
	t.ps3.SetPosition(-1, 0, 50)
	t.ps3.speed = -30.00
	a.Scene().Add(t.ps3)

	t.ps4 = NewPlayerSphere(a, "engine.ogg", &math32.Color{0, 1, 1})
	t.ps4.SetPosition(1, 0, -50)
	t.ps4.speed = 20.00
	a.Scene().Add(t.ps4)

	t.ps5 = NewPlayerSphere(a, "tone_1khz.wav", &math32.Color{1, 0, 1})
	t.ps5.SetPosition(2, 0, -50)
	t.ps5.speed = 25.00
	a.Scene().Add(t.ps5)

	t.ps6 = NewPlayerSphere(a, "tone_2khz.wav", &math32.Color{1, 1, 1})
	t.ps6.SetPosition(2, 0, -50)
	t.ps6.speed = 30.00
	a.Scene().Add(t.ps6)

	// Add controls
	if a.ControlFolder() == nil {
		return
	}
	g := a.ControlFolder().AddGroup("Play sources")
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

// Update is called every frame.
func (t *AudioDoppler) Update(a *app.App, deltaTime time.Duration) {

	t.ps1.UpdateVel(a, deltaTime)
	t.ps2.UpdateVel(a, deltaTime)
	t.ps3.UpdateVel(a, deltaTime)
	t.ps4.UpdateVel(a, deltaTime)
	t.ps5.UpdateVel(a, deltaTime)
	t.ps6.UpdateVel(a, deltaTime)
}

// Cleanup is called once at the end of the demo.
func (t *AudioDoppler) Cleanup(a *app.App) {}

func (ps *PlayerSphere) UpdateVel(a *app.App, deltaTime time.Duration) {

	delta := float32(deltaTime.Seconds()) * ps.speed
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

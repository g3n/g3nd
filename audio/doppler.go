package audio

import (
	"github.com/g3n/engine/audio"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/demos"
	"github.com/g3n/g3nd/g3nd"
)

func init() {
	demos.Map["audio.doppler"] = &AudioDoppler{}
}

type AudioDoppler struct {
	ps1 *PlayerSphere
	ps2 *PlayerSphere
	ps3 *PlayerSphere
	ps4 *PlayerSphere
	ps5 *PlayerSphere
	ps6 *PlayerSphere
}

func (t *AudioDoppler) Initialize(app *g3nd.App) {

	// Show axis helper
	ah := graphic.NewAxisHelper(1.0)
	app.Scene().Add(ah)

	// Show grid helper
	grid := graphic.NewGridHelper(100, 1, &math32.Color{0.4, 0.4, 0.4})
	app.Scene().Add(grid)

	// Sets camera position
	app.Camera().GetCamera().SetPosition(0, 4, 12)

	// Creates listener and adds it to the current camera
	listener := audio.NewListener()
	app.Camera().GetCamera().Add(listener)

	// Creates player sphere
	t.ps1 = NewPlayerSphere(app, "engine.ogg", &math32.Color{1, 0, 0})
	t.ps1.SetPosition(-3, 0, 50)
	t.ps1.speed = -20.00
	t.ps1.player.SetRolloffFactor(1)
	app.Scene().Add(t.ps1)

	t.ps2 = NewPlayerSphere(app, "tone_1khz.wav", &math32.Color{0, 1, 0})
	t.ps2.SetPosition(-2, 0, 50)
	t.ps2.speed = -25.00
	app.Scene().Add(t.ps2)

	t.ps3 = NewPlayerSphere(app, "tone_2khz.wav", &math32.Color{0, 0, 1})
	t.ps3.SetPosition(-1, 0, 50)
	t.ps3.speed = -30.00
	app.Scene().Add(t.ps3)

	t.ps4 = NewPlayerSphere(app, "engine.ogg", &math32.Color{0, 1, 1})
	t.ps4.SetPosition(1, 0, -50)
	t.ps4.speed = 20.00
	app.Scene().Add(t.ps4)

	t.ps5 = NewPlayerSphere(app, "tone_1khz.wav", &math32.Color{1, 0, 1})
	t.ps5.SetPosition(2, 0, -50)
	t.ps5.speed = 25.00
	app.Scene().Add(t.ps5)

	t.ps6 = NewPlayerSphere(app, "tone_2khz.wav", &math32.Color{1, 1, 1})
	t.ps6.SetPosition(2, 0, -50)
	t.ps6.speed = 30.00
	app.Scene().Add(t.ps6)

	// Add controls
	if app.ControlFolder() == nil {
		return
	}
	g := app.ControlFolder().AddGroup("Play sources")
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

func (t *AudioDoppler) Render(app *g3nd.App) {

	t.ps1.UpdateVel(app)
	t.ps2.UpdateVel(app)
	t.ps3.UpdateVel(app)
	t.ps4.UpdateVel(app)
	t.ps5.UpdateVel(app)
	t.ps6.UpdateVel(app)
}

func (ps *PlayerSphere) UpdateVel(app *g3nd.App) {

	delta := app.FrameDeltaSeconds() * ps.speed
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

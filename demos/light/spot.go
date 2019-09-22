package light

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/util/helper"
	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/util"
	"time"
)

func init() {
	app.DemoMap["light.spot"] = &SpotLight{}
}

type SpotLight struct {
	spot1 *util.SpotLightMesh
	spot2 *util.SpotLightMesh
	spot3 *util.SpotLightMesh
	rot   float32
}

// Start is called once at the start of the demo.
func (t *SpotLight) Start(a *app.App) {

	// Create axes helper
	axes := helper.NewAxes(1)
	a.Scene().Add(axes)

	// Sets camera position
	a.Camera().SetPosition(0, 6, 10)
	a.Camera().LookAt(&math32.Vector3{0, 0, 0}, &math32.Vector3{0, 1, 0})

	// Create box
	boxGeom := geometry.NewCube(10)
	boxMat := material.NewStandard(&math32.Color{1, 1, 1})
	boxMat.SetSide(material.SideBack)
	a.Scene().Add(graphic.NewMesh(boxGeom, boxMat))

	// Creates red spot light
	t.spot1 = util.NewSpotLightMesh(&math32.Color{1, 0, 0})
	t.spot1.Mesh.SetPosition(-1, 1, 1)
	a.Scene().Add(t.spot1)

	// Creates green spot light
	t.spot2 = util.NewSpotLightMesh(&math32.Color{0, 1, 0})
	t.spot2.Mesh.SetPosition(1, 1, -1)
	a.Scene().Add(t.spot2.Mesh)

	// Creates blue spot light
	t.spot3 = util.NewSpotLightMesh(&math32.Color{0, 0, 1})
	t.spot3.Mesh.SetPosition(0, 1, 0)
	a.Scene().Add(t.spot3.Mesh)

	// Add controls
	if a.ControlFolder() == nil {
		return
	}
	g := a.ControlFolder().AddGroup("Show lights")
	cb1 := g.AddCheckBox("Red").SetValue(t.spot1.Mesh.Visible())
	cb1.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		t.spot1.Mesh.SetVisible(!t.spot1.Mesh.Visible())
	})
	cb2 := g.AddCheckBox("Green").SetValue(t.spot2.Mesh.Visible())
	cb2.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		t.spot2.Mesh.SetVisible(!t.spot2.Mesh.Visible())
	})
	cb3 := g.AddCheckBox("Blue").SetValue(t.spot3.Mesh.Visible())
	cb3.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		t.spot3.Mesh.SetVisible(!t.spot3.Mesh.Visible())
	})
}

// Update is called every frame.
func (t *SpotLight) Update(a *app.App, deltaTime time.Duration) {

	t.rot += float32(deltaTime.Seconds())
	t.spot1.SetRotationZ(t.rot)
	t.spot2.SetRotationZ(-t.rot)
	t.spot3.SetPosition(0, 3+1.5*math32.Sin(t.rot), 0)
}

// Cleanup is called once at the end of the demo.
func (t *SpotLight) Cleanup(a *app.App) {}

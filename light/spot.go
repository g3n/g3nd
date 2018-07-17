package light

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/demos"
	"github.com/g3n/g3nd/util"
)

type SpotLight struct {
	spot1 *util.SpotLightMesh
	spot2 *util.SpotLightMesh
	spot3 *util.SpotLightMesh
	rot   float32
}

func init() {
	demos.Map["light.spot"] = &SpotLight{}
}

func (t *SpotLight) Initialize(a *app.App) {

	// Adds axis helper
	axis := graphic.NewAxisHelper(1)
	a.Scene().Add(axis)

	// Sets camera position
	a.Camera().GetCamera().SetPosition(0, 6, 10)
	a.Camera().GetCamera().LookAt(&math32.Vector3{0,0,0})

	// Creates base plane
	geom1 := geometry.NewPlane(6, 6, 16, 16)
	mat1 := material.NewPhong(&math32.Color{1, 1, 1})
	mat1.SetSide(material.SideDouble)
	plane1 := graphic.NewMesh(geom1, mat1)
	plane1.SetRotationX(math32.Pi / 2)
	a.Scene().Add(plane1)

	// Creates left plane
	geom2 := geometry.NewPlane(6, 6, 16, 16)
	mat2 := material.NewPhong(&math32.Color{1, 1, 1})
	mat2.SetSide(material.SideFront)
	plane2 := graphic.NewMesh(geom2, mat2)
	plane2.SetRotationY(math32.Pi / 2)
	plane2.SetPosition(-3, 3, 0)
	a.Scene().Add(plane2)

	// Creates right plane
	geom3 := geometry.NewPlane(6, 6, 16, 16)
	mat3 := material.NewPhong(&math32.Color{1, 1, 1})
	mat3.SetSide(material.SideFront)
	mat3.SetSpecularColor(&math32.Color{1, 1, 1})
	plane3 := graphic.NewMesh(geom3, mat3)
	plane3.SetRotationY(-math32.Pi / 2)
	plane3.SetPosition(3, 3, 0)
	a.Scene().Add(plane3)

	// Creates red spot light
	t.spot1 = util.NewSpotLightMesh(&math32.Color{1, 0, 0})
	t.spot1.Mesh.SetPosition(-1, 3, 1)
	a.Scene().Add(t.spot1)

	// Creates green spot light
	t.spot2 = util.NewSpotLightMesh(&math32.Color{0, 1, 0})
	t.spot2.Mesh.SetPosition(1, 3, -1)
	a.Scene().Add(t.spot2.Mesh)

	// Creates blue spot light
	t.spot3 = util.NewSpotLightMesh(&math32.Color{0, 0, 1})
	t.spot3.Mesh.SetPosition(0, 3, 0)
	a.Scene().Add(t.spot3.Mesh)

	// Subscribe to key events
	//	app.Gl.Subscribe(gls.OnKeyDown, t.onKey)
	//
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

func (t *SpotLight) Render(a *app.App) {

	t.rot += a.FrameDeltaSeconds()
	t.spot1.SetRotationZ(t.rot)
	t.spot2.SetRotationZ(-t.rot)
	t.spot3.SetPosition(0, 3+1.5*math32.Sin(t.rot), 0)
}


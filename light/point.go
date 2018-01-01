package light

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/demos"
	"github.com/g3n/g3nd/g3nd"

	"math"
)

type PointLight struct {
	vl    *LightMesh
	hl    *LightMesh
	count float64
}

func init() {
	demos.Map["light.point"] = &PointLight{}
}

func (t *PointLight) Initialize(app *g3nd.App) {

	// Creates right sphere
	geom1 := geometry.NewSphere(0.5, 32, 32, 0, math.Pi*2, 0, math.Pi)
	mat1 := material.NewStandard(&math32.Color{0, 0, 0.6})
	sphere1 := graphic.NewMesh(geom1, mat1)
	sphere1.SetPositionX(1)
	app.Scene().Add(sphere1)

	// Creates left sphere
	geom2 := geometry.NewSphere(0.5, 32, 32, 0, math.Pi*2, 0, math.Pi)
	mat2 := material.NewPhong(&math32.Color{0, 0.5, 0.0})
	sphere2 := graphic.NewMesh(geom2, mat2)
	sphere2.SetPositionX(-1)
	app.Scene().Add(sphere2)

	// Creates left plane
	geom3 := geometry.NewPlane(4, 4, 8, 8)
	mat3 := material.NewPhong(&math32.Color{1, 1, 1})
	pleft := graphic.NewMesh(geom3, mat3)
	pleft.SetPosition(-2, 0, 0)
	pleft.SetRotationY(math.Pi / 2)
	app.Scene().Add(pleft)

	// Creates right plane
	geom4 := geometry.NewPlane(4, 4, 8, 8)
	mat4 := material.NewStandard(&math32.Color{1, 1, 1})
	pright := graphic.NewMesh(geom4, mat4)
	pright.SetPosition(2, 0, 0)
	pright.SetRotationY(-math.Pi / 2)
	app.Scene().Add(pright)

	// Creates top plane
	geom5 := geometry.NewPlane(4, 4, 8, 8)
	mat5 := material.NewStandard(&math32.Color{1, 1, 1})
	ptop := graphic.NewMesh(geom5, mat5)
	ptop.SetPosition(0, 2, 0)
	ptop.SetRotationX(math.Pi / 2)
	app.Scene().Add(ptop)

	// Creates bottom plane
	geom6 := geometry.NewPlane(4, 4, 8, 8)
	mat6 := material.NewStandard(&math32.Color{1, 1, 1})
	pbot := graphic.NewMesh(geom6, mat6)
	pbot.SetPosition(0, -2, 0)
	pbot.SetRotationX(-math.Pi / 2)
	app.Scene().Add(pbot)

	// Creates back plane
	geom7 := geometry.NewPlane(4, 4, 8, 8)
	mat7 := material.NewStandard(&math32.Color{1, 1, 1})
	pback := graphic.NewMesh(geom7, mat7)
	pback.SetPosition(0, 0, -2)
	app.Scene().Add(pback)

	axis := graphic.NewAxisHelper(1)
	app.Scene().Add(axis)

	// Creates vertical point light
	t.vl = NewLightMesh(&math32.Color{1, 1, 1})
	t.vl.AddScene(app)

	// Creates horizontal point light
	t.hl = NewLightMesh(&math32.Color{1, 1, 1})
	t.hl.AddScene(app)

	// Add controls
	if app.ControlFolder() == nil {
		return
	}
	g := app.ControlFolder().AddGroup("Show lights")
	cb1 := g.AddCheckBox("Horizontal").SetValue(t.hl.Mesh.Visible())
	cb1.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		t.hl.Mesh.SetVisible(!t.hl.Mesh.Visible())
	})
	cb2 := g.AddCheckBox("Vertical").SetValue(t.vl.Mesh.Visible())
	cb2.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		t.vl.Mesh.SetVisible(!t.vl.Mesh.Visible())
	})
}

func (t *PointLight) Render(app *g3nd.App) {

	t.vl.Position(0, 1.5*float32(math.Sin(t.count)), 0)
	t.hl.Position(1.5*float32(math.Sin(t.count)), 1, 0)
	t.count += 0.02
}

type LightMesh struct {
	Mesh  *graphic.Mesh
	Light *light.Point
}

func NewLightMesh(color *math32.Color) *LightMesh {

	l := new(LightMesh)

	geom := geometry.NewSphere(0.05, 32, 32, 0, math.Pi*2, 0, math.Pi)
	mat := material.NewStandard(color)
	mat.SetUseLights(0)
	mat.SetEmissiveColor(color)
	l.Mesh = graphic.NewMesh(geom, mat)
	l.Mesh.SetVisible(true)

	l.Light = light.NewPoint(color, 2.0)
	l.Light.SetPosition(0, 0, 0)
	l.Light.SetLinearDecay(1)
	l.Light.SetQuadraticDecay(1)
	l.Light.SetVisible(true)

	l.Mesh.Add(l.Light)

	return l
}

func (l *LightMesh) AddScene(app *g3nd.App) {

	app.Scene().Add(l.Mesh)
}

func (l *LightMesh) Position(x, y, z float32) {

	l.Mesh.SetPosition(x, y, z)
}

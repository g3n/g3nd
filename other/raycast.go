package other

import (
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/window"
	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/demos"

	"math/rand"
)

type Raycast struct {
	rc *core.Raycaster
}

func init() {
	demos.Map["other.raycast"] = &Raycast{}
}

func (t *Raycast) Initialize(a *app.App) {

	axis := graphic.NewAxisHelper(1)
	a.Scene().Add(axis)

	l1 := light.NewDirectional(&math32.Color{1, 1, 1}, 1.0)
	l1.SetPosition(0, 0, 5)
	a.Scene().Add(l1)

	// Plane
	geom1 := geometry.NewPlane(1.5, 1, 1, 1)
	mat1 := material.NewStandard(&math32.Color{0, 1, 0})
	mat1.SetSide(material.SideFront)
	mesh1 := graphic.NewMesh(geom1, mat1)
	mesh1.SetPosition(-1.2, 0, 0)
	a.Scene().Add(mesh1)

	// Box
	geom2 := geometry.NewCube(1)
	mat2 := material.NewPhong(&math32.Color{1, 0, 0})
	mat2.SetSide(material.SideFront)
	mesh2 := graphic.NewMesh(geom2, mat2)
	mesh2.SetPosition(1.2, 0, 0)
	a.Scene().Add(mesh2)

	// Sphere
	geom3 := geometry.NewSphere(0.5, 16, 16, 0, math32.Pi*2, 0, math32.Pi)
	mat3 := material.NewStandard(&math32.Color{0, 1, 1})
	mesh3 := graphic.NewMesh(geom3, mat3)
	mesh3.SetPosition(0, 1, -1)
	a.Scene().Add(mesh3)

	// Open ended cylinder
	geom4 := geometry.NewCylinder(0.5, 0.5, 1, 16, 1, 0, 2*math32.Pi, false, false)
	mat4 := material.NewPhong(&math32.Color{1, 1, 0})
	mat4.SetSide(material.SideDouble)
	mesh4 := graphic.NewMesh(geom4, mat4)
	mesh4.SetPosition(0, -1.2, -0.5)
	a.Scene().Add(mesh4)

	// Circle
	geom5 := geometry.NewCircle(0.6, 5)
	mat5 := material.NewStandard(&math32.Color{0.5, 0.5, 0.9})
	mat5.SetSide(material.SideDouble)
	mesh5 := graphic.NewMesh(geom5, mat5)
	mesh5.SetPosition(-1.2, -1.2, -0.5)
	mesh5.SetRotation(math32.Pi/4, 0, 0)
	a.Scene().Add(mesh5)

	// Torus
	geom6 := geometry.NewTorus(0.5, 0.2, 16, 16, math32.Pi)
	mat6 := material.NewStandard(&math32.Color{0, 0, 0.5})
	mat6.SetSide(material.SideDouble)
	mesh6 := graphic.NewMesh(geom6, mat6)
	mesh6.SetPosition(1.5, -1.2, -1)
	a.Scene().Add(mesh6)

	// Cone (Cylinder)
	geom7 := geometry.NewCylinder(0, 0.5, 1, 16, 16, 0, 2*math32.Pi, true, true)
	mat7 := material.NewPhong(&math32.Color{0.8, 0.7, 0.3})
	mat7.SetSide(material.SideFront)
	mat7.SetOpacity(0.6)
	mesh7 := graphic.NewMesh(geom7, mat7)
	mesh7.SetPosition(0, 0, 0)
	a.Scene().Add(mesh7)

	// Sprite
	mat8 := material.NewStandard(&math32.Color{0, 0.3, 1})
	mesh8 := graphic.NewSprite(1, 1, mat8)
	mesh8.SetPosition(2, -2, -2)
	mesh8.SetRotationZ(math32.Pi / 4)
	mesh8.SetScale(2, 1, 1)
	a.Scene().Add(mesh8)

	// Line strip
	geom9 := geometry.NewGeometry()
	positions := math32.NewArrayF32(0, 0)
	positions.Append(
		-1, 0, -1, 1, 0, -1,
		-1, 1, -1, 1, 1, -1,
		-1, 2, -1, 1, 2, -1,
	)
	geom9.AddVBO(gls.NewVBO().AddAttrib("VertexPosition", 3).SetBuffer(positions))
	mat9 := material.NewStandard(&math32.Color{1, 0, 0})
	mesh9 := graphic.NewLineStrip(geom9, mat9)
	mesh9.SetPosition(-1.5, 0.5, -0.4)
	a.Scene().Add(mesh9)

	// Line segments
	geom10 := geometry.NewGeometry()
	positions = math32.NewArrayF32(0, 0)
	positions.Append(
		0, 0, 0, 1, 0, 0,
		0, 0, 0, -1, 0, 0,
		0, 0, 0, 0, 1, 0,
		0, 0, 0, 0, -1, 0,
		0, 0, 0, 0, 0, -1,
		0, 0, 0, 0, 0, -1,
		0, 0, 0, 0, 0, 1,
		0.1, 0.1, 0.1, 0.5, 0.5, 0.5,
	)
	geom10.AddVBO(gls.NewVBO().AddAttrib("VertexPosition", 3).SetBuffer(positions))
	mat10 := material.NewStandard(&math32.Color{0, 0, 1})
	mesh10 := graphic.NewLines(geom10, mat10)
	mesh10.SetScale(0.8, 0.8, 0.8)
	mesh10.SetPosition(1, 1.5, 0)
	a.Scene().Add(mesh10)

	// Points
	geom11 := geometry.NewGeometry()
	positions = math32.NewArrayF32(0, 0)
	for i := 0; i < 30; i++ {
		x := rand.Float32()
		y := rand.Float32()
		z := rand.Float32()
		positions.Append(x, y, z)
	}
	geom11.AddVBO(gls.NewVBO().AddAttrib("VertexPosition", 3).SetBuffer(positions))
	mat11 := material.NewPoint(&math32.Color{0, 0, 0})
	mat11.SetSize(1000)
	mesh11 := graphic.NewPoints(geom11, mat11)
	mesh11.SetPosition(-2, -1, 0)
	a.Scene().Add(mesh11)

	// Creates the raycaster
	t.rc = core.NewRaycaster(&math32.Vector3{}, &math32.Vector3{})
	t.rc.LinePrecision = 0.05
	t.rc.PointPrecision = 0.05

	// Subscribe to mouse button down events
	a.Window().Subscribe(window.OnMouseDown, func(evname string, ev interface{}) {
		t.onMouse(a, ev)
	})
}

func (t *Raycast) onMouse(a *app.App, ev interface{}) {

	// Convert mouse coordinates to normalized device coordinates
	mev := ev.(*window.MouseEvent)
	width, height := a.Window().Size()
	x := 2*(mev.Xpos/float32(width)) - 1
	y := -2*(mev.Ypos/float32(height)) + 1

	// Set the raycaster from the current camera and mouse coordinates
	a.Camera().SetRaycaster(t.rc, x, y)
	//fmt.Printf("rc:%+v\n", t.rc.Ray)

	// Checks intersection with all objects in the scene
	intersects := t.rc.IntersectObjects(a.Scene().Children(), true)
	//fmt.Printf("intersects:%+v\n", intersects)
	if len(intersects) == 0 {
		return
	}

	// Get first intersection
	obj := intersects[0].Object
	// Convert INode to IGraphic
	ig, ok := obj.(graphic.IGraphic)
	if !ok {
		a.Log().Debug("Not graphic:%T", obj)
		return
	}
	// Get graphic object
	gr := ig.GetGraphic()
	imat := gr.GetMaterial(0)

	type matI interface {
		EmissiveColor() math32.Color
		SetEmissiveColor(*math32.Color)
	}

	if v, ok := imat.(matI); ok {
		if em := v.EmissiveColor(); em.R == 1 && em.G == 1 && em.B == 1 {
			v.SetEmissiveColor(&math32.Color{0, 0, 0})
		} else {
			v.SetEmissiveColor(&math32.Color{1, 1, 1})
		}
	}

}

func (t *Raycast) Render(a *app.App) {
}

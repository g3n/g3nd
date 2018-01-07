package other

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/window"
	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/demos"

	"math"
)

func init() {
	demos.Map["other.pitch"] = &Pitch{}
}

type Pitch struct {
	base *graphic.Mesh
}

const otherPitchHelp = `
SW keys control pitch
AD keys control heading (yaw)
ZX keys control banking (roll)
R key resets to original position
`

func (t *Pitch) Initialize(a *app.App) {

	// Subscribe to key events
	a.Window().Subscribe(window.OnKeyRepeat, t.onKey)
	a.Window().Subscribe(window.OnKeyDown, t.onKey)

	// Add help label
	label1 := gui.NewLabel(otherPitchHelp)
	label1.SetFontSize(16)
	label1.SetPosition(10, 10)
	a.GuiPanel().Add(label1)

	// Top directional light
	l1 := light.NewDirectional(&math32.Color{1, 1, 1}, 0.5)
	l1.SetPosition(0, 1, 0)
	a.Scene().Add(l1)

	// Creates plane base mesh
	base_geom := geometry.NewCircle(1, 3, 0, 2*math.Pi)
	base_mat := material.NewStandard(&math32.Color{0, 1, 0})
	base_mat.SetWireframe(false)
	base_mat.SetSide(material.SideDouble)
	t.base = graphic.NewMesh(base_geom, base_mat)

	vert_geom := geometry.NewGeometry()
	positions := math32.NewArrayF32(0, 0)
	normals := math32.NewArrayF32(0, 0)
	indices := math32.NewArrayU32(0, 0)
	positions.Append(0, 0, 0, 1, 0, 0, 0, 1, 0)
	normals.Append(0, 0, 1, 0, 0, 1, 0, 0, 1)
	indices.Append(0, 1, 2)

	vert_geom.AddVBO(gls.NewVBO().AddAttrib("VertexPosition", 3).SetBuffer(positions))
	vert_geom.AddVBO(gls.NewVBO().AddAttrib("VertexNormal", 3).SetBuffer(normals))
	vert_geom.SetIndices(indices)

	vert_mat := material.NewStandard(&math32.Color{0, 0, 1})
	vert_mat.SetSide(material.SideDouble)
	vert := graphic.NewMesh(vert_geom, vert_mat)
	vert.SetScale(1.5, 1, 1)
	vert.SetPosition(-0.5, 0, 0)
	vert.SetRotationX(math.Pi / 2)
	t.base.Add(vert)

	t.base.SetScale(3, 1, 1)
	t.base.SetRotationX(-math.Pi / 2)
	t.base.SetPosition(0, 0, 0)

	a.Scene().Add(t.base)

	cam := a.Camera().GetCamera()
	cam.SetPosition(-3, 3, 3)
	spos := a.Scene().Position()
	cam.LookAt(&spos)

	// Show axis helper
	axis := graphic.NewAxisHelper(3)
	a.Scene().Add(axis)
}

func (t *Pitch) Render(a *app.App) {
}

func (t *Pitch) onKey(evname string, ev interface{}) {

	kev := ev.(*window.KeyEvent)

	var q math32.Quaternion
	xaxis := math32.Vector3{1, 0, 0}
	yaxis := math32.Vector3{0, 1, 0}
	zaxis := math32.Vector3{0, 0, 1}
	step := float32(0.01)

	// Pitch up
	if kev.Keycode == window.KeyW {
		q.SetFromAxisAngle(&yaxis, -step)
		t.base.QuaternionMult(&q)
	}
	// Pitch down
	if kev.Keycode == window.KeyS {
		q.SetFromAxisAngle(&yaxis, step)
		t.base.QuaternionMult(&q)
	}
	// Heading left
	if kev.Keycode == window.KeyA {
		q.SetFromAxisAngle(&zaxis, step)
		t.base.QuaternionMult(&q)
	}
	// Heading right
	if kev.Keycode == window.KeyD {
		q.SetFromAxisAngle(&zaxis, -step)
		t.base.QuaternionMult(&q)
	}
	// Bank left
	if kev.Keycode == window.KeyZ {
		q.SetFromAxisAngle(&xaxis, -step)
		t.base.QuaternionMult(&q)
	}
	// Bank right
	if kev.Keycode == window.KeyX {
		q.SetFromAxisAngle(&xaxis, step)
		t.base.QuaternionMult(&q)
	}
	// Reset
	if kev.Keycode == window.KeyR {
		t.base.SetRotation(-math.Pi/2, 0, 0)
	}
}

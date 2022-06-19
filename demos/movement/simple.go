package movement

/*
Gopher model was derived from the same model used in [gokoban](https://github.com/danaugrs/gokoban), which
gives the following link:

Gopher 3D model from:
https://github.com/StickmanVentures/go-gopher-model

For the purposes of this demo the model was changed in Blender
to get the origins to geometry, parented everything to body, changed
the orientation to be correct (face looking down negative Y axis).
*/

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/g3n/engine/core"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/loader/gltf"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
	"github.com/g3n/engine/util/helper"
	"github.com/g3n/engine/window"
	"github.com/g3n/g3nd/app"
)

func init() {
	app.DemoMap["movement.simple"] = &MoveSimple{}
}

type MoveSimple struct {
	a      *app.App
	gopher *core.Node

	vecVelocity, vecVelocityPaused math32.Vector3
	vecRotation, vecRotationPaused math32.Vector3

	sphere, cube *graphic.Mesh
	labelInfo    *gui.Label
	glb          *gltf.GLTF
	//this becomes a pointer to INode, this then assigned to gopher
	gophermesh core.INode
}

// Start is called once at the start of the demo.
func (t *MoveSimple) Start(a *app.App) {

	t.a = a

	// Add directional white light
	l1 := light.NewDirectional(math32.NewColor("white"), 1.0)
	l1.SetPosition(10, 10, 10)
	a.Scene().Add(l1)

	// Show grid helper
	grid := helper.NewGrid(100, 1, math32.NewColor("white"))
	a.Scene().Add(grid)

	// Add type label
	label2 := gui.NewLabel("Simple translation/rotation: uses the global axes")
	label2.SetFontSize(16)
	label2.SetColor(math32.NewColor("blue"))
	label2.SetPosition(10, 10)
	a.DemoPanel().Add(label2)

	//add info label
	t.labelInfo = gui.NewLabel("<info>")
	t.labelInfo.SetFontSize(20)
	t.labelInfo.SetColor(math32.NewColor("black"))
	t.labelInfo.SetPosition(10, 40)
	a.DemoPanel().Add(t.labelInfo)

	// Add help label
	label1 := gui.NewLabel(helpSimple)
	label1.SetFontSize(16)
	label1.SetColor(math32.NewColor("black"))
	label1.SetPosition(200, 40)
	a.DemoPanel().Add(label1)

	//load the objects
	t.loadChars(a)

	// Subscribe to key events
	a.Subscribe(window.OnKeyDown, t.onKey)

	//put everything in place
	t.doReset(a)

}

// Update is called every frame.
func (t *MoveSimple) Update(a *app.App, deltaTime time.Duration) {

	t.getCurrentInfo()
	usePos = t.gopher.Position()
	t.gopher.SetPositionVec(usePos.Add(&t.vecVelocity))
	t.gopher.RotateX(t.vecRotation.X)
	t.gopher.RotateY(t.vecRotation.Y)
	t.gopher.RotateZ(t.vecRotation.Z)
}

// Cleanup is called once at the end of the demo.
func (t *MoveSimple) Cleanup(a *app.App) {}

// Process key events
func (t *MoveSimple) onKey(evname string, ev interface{}) {

	kev := ev.(*window.KeyEvent)

	//Shift key handles rotations, sent to rotation routine
	if kev.Mods&window.ModShift > 0 {
		t.ChangeRotation(kev)
		return
	}

	incLinear = incrementLinear
	locAcceleration := incAcceleration

	//Control Key decrements velocity and acceleration
	if kev.Mods == window.ModControl {
		incLinear *= -1
		locAcceleration = 1 / locAcceleration
	}

	switch kev.Key {

	case window.KeyB: //stop all rotations
		t.vecRotation.Zero()
		t.vecRotationPaused.Zero()
		//t.vecRotationGoal.Zero()

	case window.Key0, window.KeyKP0, window.KeyO: //reset
		t.doReset(t.a)

	case window.KeyS: //stop all motion
		t.stop()

	case window.KeyT: //toggle on/off
		t.togglePause()

	case window.KeyW:
		t.vecVelocity.MultiplyScalar(locAcceleration)

		//these are local xyz/s, they move along the world axes. Rotations do not affect the linear
		//movement so good for modeling/moving satellites, thrown bottles, etc.
	case window.KeyX:
		t.vecVelocity.X += incLinear

	case window.KeyY:
		t.vecVelocity.Y += incLinear

	case window.KeyZ:
		t.vecVelocity.Z += incLinear
	}
}

//this sets the rotations that will be used in simple translation routine, called from onKey
func (t *MoveSimple) ChangeRotation(kev *window.KeyEvent) {

	incRot = incrementRotTranslate

	//Control Key decrements
	if kev.Mods&window.ModControl > 0 {
		incRot *= -1
	}

	switch kev.Key {

	case window.KeyX:
		t.vecRotation.X += incRot

	case window.KeyY:
		t.vecRotation.Y += incRot

	case window.KeyZ:
		t.vecRotation.Z += incRot
	}
}

//Stop all movements, quat slerps in go routines will finish however,
//once stopped you can continue on by pressing the movement keys again
func (t *MoveSimple) stop() {
	t.vecRotation.Zero()
	t.vecRotationPaused.Zero()

	t.vecVelocity.Zero()
	t.vecVelocityPaused.Zero()
}

//Re-set all objects to start values
func (t *MoveSimple) doReset(a *app.App) {

	t.stop()

	t.gopher.SetRotationVec(&zeroVector)
	t.gopher.SetPosition(0, 0, 0)

	a.Camera().SetRotationVec(&zeroVector)
	a.Camera().SetPositionVec(&cameraVector)

	//updates orthographic camera mode
	pos := a.Camera().Position()
	a.Camera().UpdateSize(pos.Length())

	a.Camera().LookAt(&zeroVector, math32.NewVector3(0, 1, 0))
}

func (t *MoveSimple) loadChars(a *app.App) {

	fpath := filepath.Join(a.DirData(), "gltf/Gopher/mvgopher.glb")

	ext := filepath.Ext(fpath)
	var err error

	if strings.ToUpper(ext) != ".GLB" {
		panic("movement demo Currently only supporting gltf .glb file, invalid: " + fpath)
	}

	t.glb, err = gltf.ParseBin(fpath)
	if err != nil {
		panic(err)
	}

	defaultSceneIdx := 0
	t.gophermesh, err = t.glb.LoadScene(defaultSceneIdx)

	if err != nil {
		panic(err)
	}

	a.Scene().Add(t.gophermesh)
	//re-assign INode to *Node
	t.gopher = t.gophermesh.GetNode()
	t.gopher.SetScale(0.6, 0.6, 0.6)
	t.gopher.SetPosition(0, 0, 0)

	//load geometries
	fpath = filepath.Join(a.DirData(), "images/checkerboard.jpg")

	tex1, err := texture.NewTexture2DFromImage(fpath)
	if err != nil {
		a.Log().Fatal("Error loading texture: %s", err)
	}
	tex1.SetWrapS(gls.REPEAT)
	tex1.SetWrapT(gls.REPEAT)
	tex1.SetRepeat(2, 2)

	// Creates sphere
	geom1 := geometry.NewSphere(1, 32, 32)
	mat1 := material.NewStandard(math32.NewColor("white"))
	mat1.AddTexture(tex1)
	t.sphere = graphic.NewMesh(geom1, mat1)
	t.sphere.SetPosition(0, 0, -10)
	a.Scene().Add(t.sphere)

	geom2 := geometry.NewCube(2)
	mat2 := material.NewStandard(math32.NewColor("white"))
	mat2.AddTexture(tex1)
	t.cube = graphic.NewMesh(geom2, mat2)
	t.cube.SetPosition(0, 0, 10)
	a.Scene().Add(t.cube)

	// Show axis helper
	ah := helper.NewAxes(10)
	a.Scene().Add(ah)

}

//Pause movement of the objects
func (t *MoveSimple) togglePause() {

	if t.vecVelocity.Equals(&zeroVector) {
		t.vecVelocity.Copy(&t.vecVelocityPaused)
	} else {
		t.vecVelocityPaused.Copy(&t.vecVelocity)
		t.vecVelocity.Zero()
	}

	if t.vecRotation.Equals(&zeroVector) {
		t.vecRotation.Copy(&t.vecRotationPaused)
	} else {
		t.vecRotationPaused.Copy(&t.vecRotation)
		t.vecRotation.Zero()
	}
}

//info message
func (t *MoveSimple) getCurrentInfo() {

	vecSphere := t.sphere.Position()
	vecCube = t.cube.Position()
	vecGopher = t.gopher.Position()
	vecS = *vecSphere.Sub(&vecGopher)
	vecC = *vecCube.Sub(&vecGopher)

	var sb strings.Builder

	//-----fast distance compare
	if vecS.LengthSq() >= vecC.LengthSq() {
		sb.WriteString("Cube is closer\n")
	} else {
		sb.WriteString("sphere is closer\n")
	}

	//let's see what magnitude velocity is...
	sb.WriteString(fmt.Sprintf("vel: %v\n", t.vecVelocity.Length()))

	t.labelInfo.SetText(sb.String())
}

const helpSimple = `X/Y/Z increment movement (Ctrl decrements)
Shift X/Y/Z increment rotation (Shift-Ctrl decrements)
W-accelerate      0-reset
T-Toggle Pause    B-stop rotations    S-stop all motion
`

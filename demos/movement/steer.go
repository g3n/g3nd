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
	app.DemoMap["movement.steer"] = &MoveSteer{}
}

type MoveSteer struct {
	a      *app.App
	gopher *core.Node

	vecVelocity                  math32.Vector3
	vecRotation, vecRotationGoal math32.Vector3
	//vecMovement is manipulated to travel in "face" direction, then applied to velocity
	vecMovement, vecMovementGoal math32.Vector3

	sphere, cube *graphic.Mesh
	labelInfo    *gui.Label
	glb          *gltf.GLTF
	//this becomes a pointer to INode, this then assigned to gopher
	gophermesh core.INode
}

// Start is called once at the start of the demo.
func (t *MoveSteer) Start(a *app.App) {

	t.a = a

	// Add directional white light
	l1 := light.NewDirectional(math32.NewColor("white"), 1.0)
	l1.SetPosition(10, 10, 10)
	a.Scene().Add(l1)

	// Show grid helper
	grid := helper.NewGrid(100, 1, math32.NewColor("white"))
	a.Scene().Add(grid)

	// Add type label
	label2 := gui.NewLabel("Steer / Fly: uses the local axes, always faces motion")
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
	label1 := gui.NewLabel(helpSteer)
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
func (t *MoveSteer) Update(a *app.App, deltaTime time.Duration) {
	t.getCurrentInfo()

	dtime = float32(deltaTime.Seconds())

	//approach() applies smooth motions
	t.vecRotation.X = Approach(t.vecRotationGoal.X, t.vecRotation.X, dtime/5)
	t.gopher.RotateX(t.vecRotation.X)
	t.vecRotation.Y = Approach(t.vecRotationGoal.Y, t.vecRotation.Y, dtime/5)
	t.gopher.RotateY(t.vecRotation.Y)
	t.vecRotation.Z = Approach(t.vecRotationGoal.Z, t.vecRotation.Z, dtime/5)
	t.gopher.RotateZ(t.vecRotation.Z)

	t.vecMovement.SetX(Approach(t.vecMovementGoal.X, t.vecMovement.X, dtime))
	t.vecMovement.SetY(Approach(t.vecMovementGoal.Y, t.vecMovement.Y, dtime))
	t.vecMovement.SetZ(Approach(t.vecMovementGoal.Z, t.vecMovement.Z, dtime))

	//general discussion of "steering" your character
	//see https://www.youtube.com/watch?v=FT7MShdqK6w&list=PLW3Zl3wyJwWOpdhYedlD-yCB7WQoHf-My&index=15

	//we need to calculate the two axes at 90 deg from the forward direction so we can apply trhust
	t.gopher.WorldDirection(&vecViewForward)

	t.gopher.WorldRotation(&vecViewUp)
	c := math32.Cos(vecViewUp.X) * math32.Cos(vecViewUp.Z)

	//thrusting/strafing calcs, to get the object's forward, up, right axes
	//BROKEN in that vertical can switch with horizontal sometimes, I (julius) don't yet know why.
	//maybe you can figure it out?
	vecViewUp.Set(
		math32.Cos(vecViewUp.Y)*c,
		math32.Sin(vecViewUp.Z),
		math32.Sin(vecViewUp.Y)*c)

	//2d games just need forward and right, Up (Y) can be gravity, just make sure char can't fall through floor
	vecViewForward.Normalize()
	vecViewUp.Normalize()
	vecViewUp.Cross(&vecViewForward)
	vecViewUp.Normalize()

	vecViewTmp.Copy(&vecViewUp) //Cross() modifies the vector so use a copy
	vecViewRight = *vecViewTmp.Cross(&vecViewForward)
	vecViewRight.Normalize()

	//apply the buffered (approach'd) movement to the vectors
	vecViewForward.MultiplyScalar(t.vecMovement.Z)
	vecViewRight.MultiplyScalar(t.vecMovement.X)
	vecViewUp.MultiplyScalar(t.vecMovement.Y)

	//build velocity vector from everything above
	t.vecVelocity = *vecViewForward.Add(&vecViewRight)
	t.vecVelocity.Add(&vecViewUp)

	//finally apply the manipulated velocity to the position, et voila: steering / flying motion
	usePos = t.gopher.Position()
	t.gopher.SetPositionVec(usePos.Add(&t.vecVelocity))

	//gravity (notice it is placed on movement not velocity, it will be applied next frame):
	//symbolically t.vecMovement = t.vecMovement + t.vecGravity * dtime;
	//g3n'd t.vecMovement.Add(t.vecGravity.MultiplyScalar(dtime))
	//since I didn't have a run and jump style demo I did not implement a gravity vector
	//above is how you would do it with a vecGravity like (0, -9.8, 0) where the -9.8
	//is earth's gravity attractive acceleration which will generally be in the Y axis but may be your Z
	//see https://www.youtube.com/watch?v=c4b9lCfSDQM&list=PLW3Zl3wyJwWOpdhYedlD-yCB7WQoHf-My&index=12
}

// Cleanup is called once at the end of the demo.
func (t *MoveSteer) Cleanup(a *app.App) {}

// Process key events
func (t *MoveSteer) onKey(evname string, ev interface{}) {

	kev := ev.(*window.KeyEvent)

	//send keys on to fly routines
	t.Fly(kev)

	switch kev.Key {

	case window.KeyB: //stop all rotations
		t.vecRotation.Zero()
		t.vecRotationGoal.Zero()

	case window.Key0, window.KeyKP0, window.KeyO: //reset
		t.doReset(t.a)

	case window.KeyS: //stop all motion
		t.stop()
	}
}

//this sets the motion vectors that will be used in flying using approach methods, called from onKey
func (t *MoveSteer) Fly(kev *window.KeyEvent) {

	incLinear = -incrementLinear
	locAcceleration := incAcceleration
	incRot = incrementRotFly

	//Control Key decrements velocity, acceleration, and rotation
	if kev.Mods == window.ModControl {
		incLinear *= -1
		locAcceleration = 1 / locAcceleration
		incRot *= -1
	}

	switch kev.Key {

	//This applies a LARGE change so that the approach() can be easily seen
	case window.KeyA:
		if incRot < 0 {
			t.vecRotationGoal.X += math32.Pi / 8
		} else {
			t.vecRotationGoal.X += math32.Pi / -8
		}

	case window.KeyP: //Pitch
		t.vecRotationGoal.X += incRot

	case window.KeyY: //Yaw
		t.vecRotationGoal.Y += incRot

	case window.KeyR: //Roll
		t.vecRotationGoal.Z += incRot

	case window.KeyZ: //Forward thrust
		t.vecMovementGoal.Z += -incLinear

	case window.KeyH: //horizontal thrust
		t.vecMovementGoal.X += incLinear

	case window.KeyV: //vertical thrust
		t.vecMovementGoal.Y += incLinear
	}
}

//Stop all movements, quat slerps in go routines will finish however,
//once stopped you can continue on by pressing the movement keys again
func (t *MoveSteer) stop() {
	t.vecRotation.Zero()
	t.vecRotationGoal.Zero()

	t.vecVelocity.Zero()

	t.vecMovement.Zero()
	t.vecMovementGoal.Zero()
}

//Re-set all objects to start values
func (t *MoveSteer) doReset(a *app.App) {

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

func (t *MoveSteer) loadChars(a *app.App) {
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

//an info message
func (t *MoveSteer) getCurrentInfo() {

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

//This does an easein/easeout for motion and rotation, use the deltatime and
//divide it to get longer ramp, multiply to get shorter ramp, this is not my (julius's)
//creation, see link:
//https://www.youtube.com/watch?v=qJq7I2DLGzI&list=PLW3Zl3wyJwWOpdhYedlD-yCB7WQoHf-My&index=13
func Approach(goal, current, dtime float32) float32 {
	flDifference = goal - current
	if flDifference > dtime {
		return current + dtime
	}
	if flDifference < -dtime {
		return current - dtime
	}
	return goal
}

const helpSteer = `Z  forward/backward thrust
P Pitch (X axis)    Y Yaw (Y axis)      R Roll (Z axis) 
H Horizontal (left/right) thrust (can flip)
V Vertical (up/down) thrust (can flip)
A Fast spin demo
(Ctrl decrements all above)
B-stop rotations    S-stop all motion    0-reset
`

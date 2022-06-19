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
	"math/rand"
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
	"github.com/g3n/engine/window"
	"github.com/g3n/g3nd/app"
)

func init() {
	app.DemoMap["movement.slerp"] = &MoveSlerp{}
}

type MoveSlerp struct {
	a      *app.App
	gopher *core.Node

	//used for slerp'ing
	fromQuat, toQuat math32.Quaternion

	sphere, cube, cone *graphic.Mesh
	glb                *gltf.GLTF
	//this becomes a pointer to INode, this then assigned to gopher
	gophermesh core.INode
}

var (
	//chooses objects in order for gopher to LookAt
	toggleLookAtTarget int = -1
	coord              float32
	//let it slerp
	slerping bool = false
	//used in quaternion slerping through a go routine
	vecLookAt, vecLookAtLooker, vecLookAtTarget math32.Vector3
	rotMatrix                                   math32.Matrix4
)

// Start is called once at the start of the demo.
func (t *MoveSlerp) Start(a *app.App) {

	t.a = a

	// Add directional white light
	l1 := light.NewDirectional(math32.NewColor("white"), 1.0)
	l1.SetPosition(10, 10, 10)
	a.Scene().Add(l1)

	// Add type label
	label2 := gui.NewLabel("LookAt's: both smooth Quaternion SLERP and 'snap-to'")
	label2.SetFontSize(16)
	label2.SetColor(math32.NewColor("blue"))
	label2.SetPosition(10, 10)
	a.DemoPanel().Add(label2)

	// Add help label
	label1 := gui.NewLabel(helpSlerp)
	label1.SetFontSize(16)
	label1.SetColor(math32.NewColor("black"))
	label1.SetPosition(10, 40)
	a.DemoPanel().Add(label1)

	a.Camera().SetPositionVec(math32.NewVector3(18, 6, -2))
	//updates orthographic camera mode
	pos := a.Camera().Position()
	a.Camera().UpdateSize(pos.Length())
	a.Camera().LookAt(math32.NewVector3(0, 0, 2), math32.NewVector3(0, 1, 0))

	//load the objects
	t.loadChars(a)

	// Subscribe to key events
	a.Subscribe(window.OnKeyDown, t.onKey)

}

// Process key events
func (t *MoveSlerp) onKey(evname string, ev interface{}) {

	kev := ev.(*window.KeyEvent)

	switch kev.Key {

	case window.Key1, window.KeyKP1:
		t.sphere.SetPositionVec(t.getRandomPositionVec())

	case window.Key2, window.KeyKP2:
		t.cube.SetPositionVec(t.getRandomPositionVec())

	case window.Key3, window.KeyKP3:
		t.cone.SetPositionVec(t.getRandomPositionVec())

	case window.Key4, window.KeyKP4:
		t.gopher.SetPositionVec(t.getRandomPositionVec())

	case window.KeyL: //LookAt's, direct and SLERP
		if slerping {
			return
		}

		t.getSlerpVector()

		//Control was pressed, use the Direct LookAt()
		if kev.Mods&window.ModControl > 0 {
			t.gopher.LookAt(&vecLookAt, math32.NewVector3(0, 1, 0))
			return
		}

		//Control Key was not pressed, use the SLERP LookAt()
		t.getSlerpQuats()
		go t.quatSlerp(30, &t.fromQuat, &t.toQuat)
	}
}

// Update is called every frame.
func (t *MoveSlerp) Update(a *app.App, deltaTime time.Duration) {}

// Cleanup is called once at the end of the demo.
func (t *MoveSlerp) Cleanup(a *app.App) {}

//gets random postion vectors in range of 3 to around 9
func (t *MoveSlerp) getRandomPositionVec() *math32.Vector3 {
	result := math32.NewVec3()

	rSource := rand.NewSource(time.Now().UnixNano())
	rGen := rand.New(rSource)

	var factor int
	getFactor := func() float32 {
		factor = rGen.Intn(100)
		if factor%2 == 0 {
			return 1
		}
		return -1
	}

	coord = (3 + (rGen.Float32() * 6)) * getFactor()
	result.SetX(coord)
	coord = (3 + (rGen.Float32() * 6)) * getFactor()
	result.SetY(coord)
	coord = (3 + (rGen.Float32() * 6)) * getFactor()
	result.SetZ(coord)

	return result
}

func (t *MoveSlerp) loadChars(a *app.App) {
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

	geom3 := geometry.NewCone(1, 1, 16, 8, true)
	mat3 := material.NewStandard(math32.NewColor("white"))
	mat3.AddTexture(tex1)
	t.cone = graphic.NewMesh(geom3, mat3)
	t.cone.SetPosition(0, 5, 0)
	a.Scene().Add(t.cone)
}

//Does the work of changing the LookAt targets and calculating the
//"fixed" vector
func (t *MoveSlerp) getSlerpVector() {

	toggleLookAtTarget++
	switch toggleLookAtTarget % 3 {
	case 0:
		t.sphere.WorldPosition(&vecLookAtTarget)
	case 1:
		t.cube.WorldPosition(&vecLookAtTarget)
	case 2:
		t.cone.WorldPosition(&vecLookAtTarget)
	}

	t.gopher.WorldPosition(&vecLookAtLooker)

	//These calcs must be done for object lookAt's, there is a disconnect
	//with object LookAt's which use, I believe, a camera LookAt, which is wrong
	//for an object. I stumbled across this vector fix. Literally. I made a mistake
	//and subtracted twice when I meant to comment one out. Wow-ouch.

	//"fixes" the vector, the subtraction order matters
	vecLookAt = *vecLookAtTarget.Sub(&vecLookAtLooker)
	vecLookAt.SubVectors(&vecLookAtLooker, &vecLookAtTarget)
}

//Does the work of getting the from and to quaternions needed for a smooth slerp
func (t *MoveSlerp) getSlerpQuats() {

	rotMatrix.LookAt(&vecLookAtLooker, &vecLookAt, math32.NewVector3(0, 1, 0))

	t.toQuat.SetFromRotationMatrix(&rotMatrix)
	t.toQuat.Normalize() //whoops! Yes, you need to do this

	t.fromQuat = t.gopher.Quaternion()
	t.fromQuat.Normalize() //whoops! Yes, you need to do this
}

//given from/to quaternion do a sherical linear interpolation, this is called as a go func,
//cnt is how many interpolations you want per given a ticker of 1/60 of a second,
//do NOT use this on an object that is being rendered in the render loop, you will not be happy
func (t *MoveSlerp) quatSlerp(cnt float32, from, to *math32.Quaternion) {

	slerping = true
	ticker := time.NewTicker(time.Millisecond * 34) //about 60 times a second

	for range ticker.C {
		//the Slerp() func changes the slerped quat. if you leave this alone the
		//slerping accelerates because the distance to rotate gets smaller and smaller.
		//this is my method to get a linear slerp, by using the inverse it adjusts for
		//the changing slerp length
		t.gopher.SetRotationQuat(from.Slerp(to, 1/cnt))
		cnt--
		if cnt <= 0 {
			ticker.Stop()
			break
		}
	}
	slerping = false
	//g.Log.Info("yeah you need the break statement")
}

const helpSlerp = `L  LookAt objects in turn (smooth, wait for slerp to finish)
Ctrl-L  LookAt objects in turn (snap-to)
Move:  1=sphere    2=cube   3=cone   4=GOpher
`

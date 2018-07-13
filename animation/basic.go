package animation

import (
	"github.com/g3n/engine/graphic"
	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/demos"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/animation"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/gui"
)

func init() {
	demos.Map["animation.basic"] = &AnimationBasic{}
}

type AnimationBasic struct {
	anim *animation.Animation
}

func (t *AnimationBasic) Initialize(a *app.App) {

	ah := graphic.NewAxisHelper(1.0)
	a.Scene().Add(ah)

	// Adds white directional front light
	dir1 := light.NewDirectional(&math32.Color{1, 1, 1}, 1.0)
	dir1.SetPosition(0, 5, 10)
	a.Scene().Add(dir1)

	// Add box
	geom := geometry.NewSegmentedCube(1,2)
	mat := material.NewStandard(&math32.Color{0.5, 0, 0})
	mat.SetWireframe(false)
	box := graphic.NewMesh(geom, mat)
	a.Scene().Add(box)

	// Create a looping animation
	t.anim = animation.NewAnimation()
	t.anim.SetLoop(true)

	keyframes := math32.NewArrayF32(0, 3)
	keyframes.Append(0, 2, 4, 6)

	posValues := math32.NewArrayF32(0, 9)
	posValues.AppendVector3(&math32.Vector3{0,0,0}, &math32.Vector3{1,0,0}, &math32.Vector3{1,1,0}, &math32.Vector3{0,0,0})

	scaleValues := math32.NewArrayF32(0, 9)
	scaleValues.AppendVector3(&math32.Vector3{1,1,1}, &math32.Vector3{0.4,0.4,0.4}, &math32.Vector3{3,0.4,2}, &math32.Vector3{1,1,1})

	posChan := animation.NewPositionChannel(box)
	posChan.SetBuffers(keyframes, posValues)
	t.anim.AddChannel(posChan)

	scaleChan := animation.NewScaleChannel(box)
	scaleChan.SetBuffers(keyframes, scaleValues)
	t.anim.AddChannel(scaleChan)

	// Add animation controls
	g1 := a.ControlFolder().AddGroup("Animation")
	cb1 := g1.AddCheckBox("Paused").SetValue(false)
	cb1.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		t.anim.SetPaused(!t.anim.Paused())
	})
	cb2 := g1.AddCheckBox("Loop").SetValue(true)
	cb2.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		t.anim.SetLoop(!t.anim.Loop())
	})
}

func (t *AnimationBasic) Render(a *app.App) {

	t.anim.Update(a.FrameDeltaSeconds())
}

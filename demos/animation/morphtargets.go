package animation

import (
	"github.com/g3n/engine/animation"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
	"time"

	"github.com/g3n/g3nd/app"
)

func init() {
	app.DemoMap["animation.morphtargets"] = &AnimationMorphTargets{}
}

type AnimationMorphTargets struct {
	morphMesh *graphic.Mesh
	morphGeom *geometry.MorphGeometry
	count     float32
	weights   []float32
	anim      *animation.Animation

	g1s1 *gui.Slider
	g1s2 *gui.Slider
}

// Start is called once at the start of the demo.
func (t *AnimationMorphTargets) Start(a *app.App) {

	// Adds white directional front light
	dir1 := light.NewDirectional(&math32.Color{1, 1, 1}, 1.0)
	dir1.SetPosition(0, 5, 10)
	a.Scene().Add(dir1)

	// Create base geometry
	geom := geometry.NewSphere(1, 16, 16, 0, math32.Pi*2, 0, math32.Pi)
	t.morphGeom = geometry.NewMorphGeometry(geom.GetGeometry())

	var target1Vertices, target2Vertices math32.ArrayF32
	vertIdx := 0
	geom.ReadVertices(func(vertex math32.Vector3) bool {
		var vertex1, vertex2 math32.Vector3
		vertex1.Copy(&vertex)
		vertex2.Copy(&vertex)
		vertex1.Y += math32.Sin(vertex1.X)
		vertex1.Z += math32.Cos(vertex1.Y)
		//vertex1.X *= 2
		vertex2.X -= math32.Sin(vertex2.Y)
		vertex2.Z -= math32.Cos(vertex2.X)
		//vertex2.Y *= 2
		target1Vertices.AppendVector3(&vertex1)
		target2Vertices.AppendVector3(&vertex2)
		vertIdx++
		return false
	})
	target1 := geometry.NewGeometry()
	target1.AddVBO(gls.NewVBO(target1Vertices).AddAttrib(gls.VertexPosition))
	target2 := geometry.NewGeometry()
	target2.AddVBO(gls.NewVBO(target2Vertices).AddAttrib(gls.VertexPosition))
	t.morphGeom.AddMorphTargets(target1, target2)

	// Create texture
	texfile := a.DirData() + "/images/checkerboard.jpg"
	tex1, err := texture.NewTexture2DFromImage(texfile)
	if err != nil {
		a.Log().Fatal("Error loading texture: %s", err)
	}
	tex1.SetWrapS(gls.REPEAT)
	tex1.SetWrapT(gls.REPEAT)
	tex1.SetRepeat(2, 2)

	mat := material.NewPhong(&math32.Color{0, 0, 1})
	mat.AddTexture(tex1)
	t.morphMesh = graphic.NewMesh(t.morphGeom, mat)
	a.Scene().Add(t.morphMesh)

	weights := []float32{0.5, 0.5}
	t.morphGeom.SetWeights(weights)

	// Add controls (read-only)
	if a.ControlFolder() == nil {
		return
	}
	g1 := a.ControlFolder().AddGroup("Morph target weights")
	t.g1s1 = g1.AddSlider("Weight 0:", 1, 0.5)
	t.g1s1.SetEnabled(false)
	t.g1s2 = g1.AddSlider("Weight 1:", 1, 0.5)
	t.g1s2.SetEnabled(false)

	// Create a looping animation
	t.anim = animation.NewAnimation()
	t.anim.SetLoop(true)

	keyframes := math32.NewArrayF32(0, 3)
	keyframes.Append(0, 2, 4, 6, 8)

	weightValues := math32.NewArrayF32(0, 9)
	weightValues.AppendVector2(
		&math32.Vector2{0, 0},
		&math32.Vector2{0, 1},
		&math32.Vector2{1, 0},
		&math32.Vector2{1, 1},
		&math32.Vector2{0, 0},
	)

	morphChan := animation.NewMorphChannel(t.morphGeom)
	morphChan.SetBuffers(keyframes, weightValues)
	t.anim.AddChannel(morphChan)
}

// Update is called every frame.
func (t *AnimationMorphTargets) Update(a *app.App, deltaTime time.Duration) {

	t.anim.Update(float32(deltaTime.Seconds()))
	updatedWeights := t.morphGeom.Weights()
	t.g1s1.SetValue(updatedWeights[0])
	t.g1s2.SetValue(updatedWeights[1])

	// Rotate at 1 rotation each 10 seconds
	delta := float32(deltaTime.Seconds()) * 2 * math32.Pi / 10
	t.morphMesh.RotateY(delta)
}

// Cleanup is called once at the end of the demo.
func (t *AnimationMorphTargets) Cleanup(a *app.App) {}

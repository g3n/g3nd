package other

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/demos"

	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/texture"
)

func init() {
	demos.Map["other.morphtargets"] = &MorphTargets{}
}

type MorphTargets struct {
	morphMesh *graphic.Mesh
	count     float32
	weights   []float32
}

func (t *MorphTargets) Initialize(a *app.App) {

	// Adds white directional front light
	dir1 := light.NewDirectional(&math32.Color{1, 1, 1}, 1.0)
	dir1.SetPosition(0, 5, 10)
	a.Scene().Add(dir1)

	// Create base geometry
	geom := geometry.NewSphere(1, 16, 16, 0, math32.Pi*2, 0, math32.Pi)
	morphGeom := geometry.NewMorphGeometry(geom.GetGeometry())

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
	morphGeom.AddMorphTargets(target1, target2)

	// Create texture
	texfile := a.DirData() + "/images/checkerboard.jpg"
	tex1, err := texture.NewTexture2DFromImage(texfile)
	if err != nil {
		a.Log().Fatal("Error loading texture: %s", err)
	}
	tex1.SetWrapS(gls.REPEAT)
	tex1.SetWrapT(gls.REPEAT)
	tex1.SetRepeat(2, 2)
	
	mat := material.NewPhong(&math32.Color{0, 1, 0})
	mat.AddTexture(tex1)
	t.morphMesh = graphic.NewMesh(morphGeom, mat)
	a.Scene().Add(t.morphMesh)

	weights := []float32{0.5, 0.5}
	morphGeom.SetWeights(weights)

	// Add controls
	if a.ControlFolder() == nil {
		return
	}
	g1 := a.ControlFolder().AddGroup("Morph target weights")
	g1s1 := g1.AddSlider("Weight 0:", 1, 0.5)
	g1s1.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		weights[0] = g1s1.Value()
		morphGeom.SetWeights(weights)
	})
	g1s2 := g1.AddSlider("Weight 1:", 1, 0.5)
	g1s2.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		weights[1] = g1s2.Value()
		morphGeom.SetWeights(weights)
	})

}

func (t *MorphTargets) Render(a *app.App) {

	// Rotate at 1 rotation each 10 seconds
	delta := a.FrameDeltaSeconds() * 2 * math32.Pi / 10
	t.morphMesh.RotateY(delta)
}

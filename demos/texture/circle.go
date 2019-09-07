package texture

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
	"github.com/g3n/g3nd/app"
	"time"
)

func init() {
	app.DemoMap["texture.circle"] = &Texcircle{}
}

type Texcircle struct {
	mesh1 *graphic.Mesh
	mesh2 *graphic.Mesh
	mesh3 *graphic.Mesh
}

// Start is called once at the start of the demo.
func (t *Texcircle) Start(a *app.App) {

	// Adds white directional front light
	dir1 := light.NewDirectional(&math32.Color{1, 1, 1}, 1.0)
	dir1.SetPosition(0, 0, 10)
	a.Scene().Add(dir1)

	geom1 := geometry.NewCircle(1, 3)
	mat1 := material.NewStandard(&math32.Color{0, 1, 0})
	mat1.SetWireframe(false)
	tex1 := texture.NewBoard(32, 32, math32.NewColor("white"), math32.NewColor("black"), math32.NewColor("black"), math32.NewColor("white"), 0.8)
	tex1.SetWrapS(gls.REPEAT)
	tex1.SetWrapT(gls.REPEAT)
	tex1.SetRepeat(4, 4)
	mat1.AddTexture(tex1)
	t.mesh1 = graphic.NewMesh(geom1, mat1)
	t.mesh1.SetVisible(true)
	t.mesh1.SetPositionX(-1.5)
	t.mesh1.SetPositionZ(0)
	a.Scene().Add(t.mesh1)

	geom2 := geometry.NewCircle(1, 50)
	mat2 := material.NewStandard(&math32.Color{0.5, 0.5, 0.5})
	tex2, err := texture.NewTexture2DFromImage(a.DirData() + "/images/tiger1.jpg")
	if err != nil {
		a.Log().Fatal("Error loading texture: %s", err)
	}
	mat2.SetSide(material.SideDouble)
	mat2.AddTexture(tex2)
	t.mesh2 = graphic.NewMesh(geom2, mat2)
	t.mesh2.SetVisible(true)
	t.mesh2.SetPositionX(0)
	t.mesh2.SetPositionZ(-0.5)
	a.Scene().Add(t.mesh2)

	geom3 := geometry.NewCircle(1, 5)
	mat3 := material.NewStandard(&math32.Color{1, 0, 0})
	tex3 := texture.NewBoard(32, 32, math32.NewColor("white"), math32.NewColor("black"), math32.NewColor("black"), math32.NewColor("white"), 0.6)
	tex3.SetWrapS(gls.REPEAT)
	tex3.SetWrapT(gls.REPEAT)
	tex3.SetRepeat(4, 4)
	mat3.AddTexture(tex3)
	t.mesh3 = graphic.NewMesh(geom3, mat3)
	t.mesh3.SetVisible(true)
	t.mesh3.SetPositionX(2.0)
	t.mesh3.SetPositionZ(-1.0)
	a.Scene().Add(t.mesh3)

	// Show axis helper
	axis := graphic.NewAxisHelper(2)
	a.Scene().Add(axis)
}

// Update is called every frame.
func (t *Texcircle) Update(a *app.App, deltaTime time.Duration) {

	// TODO use deltaTime
	t.mesh1.RotateZ(0.01)
	t.mesh3.RotateZ(-0.01)
}

// Cleanup is called once at the end of the demo.
func (t *Texcircle) Cleanup(a *app.App) {}

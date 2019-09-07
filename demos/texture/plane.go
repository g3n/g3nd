package texture

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
	"github.com/g3n/g3nd/app"
	"time"
)

func init() {
	app.DemoMap["texture.plane"] = &Texplane{}
}

type Texplane struct {
	plane1 *graphic.Mesh
	plane2 *graphic.Mesh
}

// Start is called once at the start of the demo.
func (t *Texplane) Start(a *app.App) {

	axis := graphic.NewAxisHelper(1)
	a.Scene().Add(axis)

	// Adds red directional right light
	dir1 := light.NewDirectional(&math32.Color{1, 0, 0}, 1.0)
	dir1.SetPosition(1, 0, 0)
	a.Scene().Add(dir1)

	// Adds green directional top light
	dir2 := light.NewDirectional(&math32.Color{1, 0, 0}, 1.0)
	dir2.SetPosition(0, 1, 0)
	a.Scene().Add(dir2)

	// Adds white directional front light
	dir3 := light.NewDirectional(&math32.Color{1, 1, 1}, 1.0)
	dir3.SetPosition(0, 0, 1)
	a.Scene().Add(dir3)

	// Loads texture from image
	texfile := a.DirData() + "/images/uvgrid.jpg"
	tex1, err := texture.NewTexture2DFromImage(texfile)
	if err != nil {
		a.Log().Fatal("Error:%s loading texture:%s", err, texfile)
	}
	// Creates plane 1
	plane1_geom := geometry.NewPlane(2, 2, 1, 1)
	plane1_mat := material.NewStandard(&math32.Color{1, 1, 1})
	plane1_mat.SetSide(material.SideDouble)
	plane1_mat.AddTexture(tex1)
	t.plane1 = graphic.NewMesh(plane1_geom, plane1_mat)
	t.plane1.SetPosition(0, 1.1, 0)
	a.Scene().Add(t.plane1)

	// Loads texture from image
	texfile = a.DirData() + "/images/tiger1.jpg"
	tex2, err := texture.NewTexture2DFromImage(texfile)
	if err != nil {
		a.Log().Fatal("Error:%s loading texture:%s", err, texfile)
	}
	// Creates plane2
	plane2_geom := geometry.NewPlane(2, 2, 1, 1)
	plane2_mat := material.NewPhong(&math32.Color{1, 1, 1})
	plane2_mat.SetSide(material.SideDouble)
	plane2_mat.AddTexture(tex2)
	t.plane2 = graphic.NewMesh(plane2_geom, plane2_mat)
	t.plane2.SetPosition(0, -1.1, 0)
	a.Scene().Add(t.plane2)
}

// Update is called every frame.
func (t *Texplane) Update(a *app.App, deltaTime time.Duration) {

	// TODO use deltaTime
	t.plane1.RotateY(0.01)
	t.plane2.RotateY(-0.01)
}

// Cleanup is called once at the end of the demo.
func (t *Texplane) Cleanup(a *app.App) {}

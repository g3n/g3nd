package texture

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
	"github.com/g3n/engine/window"
	"github.com/g3n/g3nd/demos"
	"github.com/g3n/g3nd/g3nd"
)

type Texbox struct {
	tex1 *texture.Texture2D
	tex2 *texture.Texture2D
	tex3 *texture.Texture2D
	mat1 *material.Standard
	mat2 *material.Phong
	mat3 *material.Standard
	mat4 *material.Phong
	box1 *graphic.Mesh
	box2 *graphic.Mesh
	box3 *graphic.Mesh
	box4 *graphic.Mesh
}

func init() {
	demos.Map["texture.box"] = &Texbox{}
}

func (t *Texbox) Initialize(app *g3nd.App) {

	axis := graphic.NewAxisHelper(1)
	app.Scene().Add(axis)

	// Adds white directional front light
	dir1 := light.NewDirectional(&math32.Color{1, 1, 1}, 1.0)
	dir1.SetPosition(0, 0, 10)
	app.Scene().Add(dir1)

	// Adds blue directional right light
	dir2 := light.NewDirectional(&math32.Color{0, 0, 1}, 1.0)
	dir2.SetPosition(10, 0, 0)
	app.Scene().Add(dir2)

	// Adds red directional left light
	dir3 := light.NewDirectional(&math32.Color{1, 0, 0}, 1.0)
	dir3.SetPosition(-10, 0, 0)
	app.Scene().Add(dir3)

	// Creates texture1
	texfile := app.DirData() + "/images/checkerboard.jpg"
	tex1, err := texture.NewTexture2DFromImage(texfile)
	if err != nil {
		app.Log().Fatal("Error:%s loading texture:%s", err, texfile)
	}
	t.tex1 = tex1
	// Creates box 1
	geom1 := geometry.NewBox(1.5, 1.5, 1.5, 16, 16, 16)
	t.mat1 = material.NewStandard(&math32.Color{1, 1, 1})
	t.mat1.AddTexture(t.tex1)
	t.box1 = graphic.NewMesh(geom1, t.mat1)
	t.box1.SetPosition(-1.5, 1, 0)
	app.Scene().Add(t.box1)

	// Creates texture2
	texfile = app.DirData() + "/images/brick1.jpg"
	tex2, err := texture.NewTexture2DFromImage(texfile)
	if err != nil {
		app.Log().Fatal("Error:%s loading texture:%s", err, texfile)
	}
	t.tex2 = tex2
	// Creates box 2
	geom2 := geometry.NewBox(1.5, 1.5, 1.5, 16, 16, 16)
	t.mat2 = material.NewPhong(&math32.Color{1, 1, 1})
	t.mat2.AddTexture(t.tex2)
	t.box2 = graphic.NewMesh(geom2, t.mat2)
	t.box2.SetPosition(1.5, 1, 0)
	app.Scene().Add(t.box2)

	// Creates texture3
	texfile = app.DirData() + "/images/moss.png"
	tex3, err := texture.NewTexture2DFromImage(texfile)
	if err != nil {
		app.Log().Fatal("Error:%s loading texture:%s", err, texfile)
	}
	t.tex3 = tex3
	// Creates box 3
	geom3 := geometry.NewBox(1.5, 1.5, 1.5, 16, 16, 16)
	t.mat3 = material.NewStandard(&math32.Color{1, 1, 1})
	t.mat3.AddTexture(t.tex1.Incref())
	t.mat3.AddTexture(t.tex3)
	t.box3 = graphic.NewMesh(geom3, t.mat3)
	t.box3.SetPosition(-1.5, -1, 0)
	app.Scene().Add(t.box3)

	// Creates box 4
	geom4 := geometry.NewBox(1.5, 1.5, 1.5, 16, 16, 16)
	t.mat4 = material.NewPhong(&math32.Color{1, 1, 1})
	t.mat4.AddTexture(t.tex2.Incref())
	t.mat4.AddTexture(t.tex3.Incref())
	t.box4 = graphic.NewMesh(geom4, t.mat4)
	t.box4.SetPosition(1.5, -1, 0)
	app.Scene().Add(t.box4)

	// Subscribe to key events
	app.Window().Subscribe(window.OnKeyDown, func(evname string, ev interface{}) {
		t.onKey(app, ev)
	})
}

func (t *Texbox) Render(app *g3nd.App) {

	t.box1.AddRotationY(0.01)
	t.box2.AddRotationY(-0.01)
	t.box3.AddRotationY(0.01)
	t.box4.AddRotationY(-0.01)
}

func (t *Texbox) onKey(app *g3nd.App, ev interface{}) {

	kev := ev.(*window.KeyEvent)
	if kev.Action == window.Release {
		return
	}
	switch kev.Keycode {
	case window.Key1:
		t.tex1.SetVisible(!t.tex1.Visible())
	case window.Key2:
		t.tex2.SetVisible(!t.tex2.Visible())
	case window.Key3:
		t.tex3.SetVisible(!t.tex3.Visible())
	case window.Key4:
		err := t.tex2.SetImage(app.DirData() + "/images/wall1.jpg")
		if err != nil {
			app.Log().Fatal("Error:%s loading texture", err)
		}
	case window.Key5:
		err := t.tex2.SetImage(app.DirData() + "/images/brick1.jpg")
		if err != nil {
			app.Log().Fatal("Error:%s loading texture", err)
		}
	case window.Key6:
		if t.mat4.HasTexture(t.tex2) {
			t.mat4.RemoveTexture(t.tex2)
		} else {
			t.mat4.AddTexture(t.tex2)
		}
	case window.Key7:
		if t.mat4.HasTexture(t.tex3) {
			t.mat4.RemoveTexture(t.tex3)
		} else {
			t.mat4.AddTexture(t.tex3)
		}
	}
}

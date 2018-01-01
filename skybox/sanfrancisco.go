package skybox

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
	"github.com/g3n/g3nd/demos"
	"github.com/g3n/g3nd/g3nd"
)

func init() {
	demos.Map["skybox.sanfrancisco"] = &Skybox{}
}

type Skybox struct {
}

func (t *Skybox) Initialize(app *g3nd.App) {

	var textures = []string{
		"sanfrancisco/posx.jpg",
		"sanfrancisco/negx.jpg",
		"sanfrancisco/posy.jpg",
		"sanfrancisco/negy.jpg",
		"sanfrancisco/posz.jpg",
		"sanfrancisco/negz.jpg",
	}

	// Add axis helper
	axis := graphic.NewAxisHelper(2)
	app.Scene().Add(axis)

	geom := geometry.NewBox(50, 50, 50, 2, 2, 2)
	skybox := graphic.NewMesh(geom, nil)
	for i := 0; i < 6; i++ {
		tex, err := texture.NewTexture2DFromImage(app.DirData() + "/images/" + textures[i])
		if err != nil {
			app.Log().Fatal("Error loading texture: %s", err)
		}
		matFace := material.NewStandard(&math32.Color{1, 1, 1})
		matFace.AddTexture(tex)
		matFace.SetSide(material.SideBack)
		skybox.AddGroupMaterial(matFace, i)
	}
	app.Scene().Add(skybox)
}

func (t *Skybox) Render(app *g3nd.App) {
}

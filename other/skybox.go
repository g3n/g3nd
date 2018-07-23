package other

import (
	"github.com/g3n/engine/graphic"
	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/demos"
)

func init() {
	demos.Map["other.skybox"] = &Skybox{}
}

type Skybox struct {
}

func (t *Skybox) Initialize(a *app.App) {

	// Create Skybox
	skybox, err := graphic.NewSkybox(graphic.SkyboxData{
		a.DirData() + "/images/sanfrancisco/", "jpg",
		[6]string{"posx", "negx", "posy", "negy", "posz", "negz"}})
	if err != nil {
		panic(err)
	}
	a.Scene().Add(skybox)

	// Add axis helper
	axis := graphic.NewAxisHelper(2)
	a.Scene().Add(axis)
}

func (t *Skybox) Render(a *app.App) {
}

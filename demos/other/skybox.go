package other

import (
	"github.com/g3n/engine/graphic"
	"github.com/g3n/g3nd/app"
	"time"
)

func init() {
	app.DemoMap["other.skybox"] = &Skybox{}
}

type Skybox struct{}

// Start is called once at the start of the demo.
func (t *Skybox) Start(a *app.App) {

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

// Update is called every frame.
func (t *Skybox) Update(a *app.App, deltaTime time.Duration) {}

// Cleanup is called once at the end of the demo.
func (t *Skybox) Cleanup(a *app.App) {}

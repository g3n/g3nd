package other

import (
	"time"

	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/util/helper"
	"github.com/g3n/g3nd/app"
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

	// Create axes helper
	axes := helper.NewAxes(2)
	a.Scene().Add(axes)
}

// Update is called every frame.
func (t *Skybox) Update(a *app.App, deltaTime time.Duration) {}

// Cleanup is called once at the end of the demo.
func (t *Skybox) Cleanup(a *app.App) {}

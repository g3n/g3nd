// This is a simple model for your tests
package tests

import (
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/demos"
	"github.com/g3n/g3nd/g3nd"
)

// Sets the category and name of your test in the demos.Map
// The category name choosen here starts with a "|" so it shows as the
// last category in list. Change "model" to the name of your test.
func init() {
	demos.Map["|tests|.model"] = &testsModel{}
}

// This is your test object. You can store state here.
// By convention and to avoid conflict with other demo/tests name it
// using your test category and name.
type testsModel struct {
	grid *graphic.GridHelper // Pointer to a GridHelper created in 'Initialize'
}

// This method will be called once when the test is selected from the G3ND list.
// app is a pointer to the G3ND application.
// It allows access to several methods such as app.Scene(), which returns the current scene,
// app.GuiPanel(), app.Camera(), app.Window() among others.
// You can build your scene adding your objects to the app.Scene()
func (t *testsModel) Initialize(app *g3nd.App) {

	// Show axis helper
	ah := graphic.NewAxisHelper(1.0)
	app.Scene().Add(ah)

	// Creates a grid helper and saves its pointer in the test state
	t.grid = graphic.NewGridHelper(50, 1, &math32.Color{0.4, 0.4, 0.4})
	app.Scene().Add(t.grid)

	// Changes the camera position
	app.Camera().GetCamera().SetPosition(0, 4, 10)
}

// This method will be called at every frame
// You can animate your objects here.
func (t *testsModel) Render(app *g3nd.App) {

	// Rotate the grid, just for show.
	rps := app.FrameDeltaSeconds() * 2 * math32.Pi
	t.grid.AddRotationY(rps * 0.05)
}

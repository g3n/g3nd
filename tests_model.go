// This is a simple model for your tests
package main

import (
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/math32"
)

// Sets the category and name of your test in the package global "TestMap"
// The category name choosen here starts with a "|" so it shows as the
// last category in list. Change "model" to the name of your test.
func init() {
	TestMap["|tests|.model"] = &testsModel{}
}

// This is your test object. You can store state here.
// This name must be unique in the package
type testsModel struct {
	grid *graphic.GridHelper
}

// This method will be called once when the test is selected from the list
// You can add your objects to the scene here.
// The ctx objects contain several global objects built by the main program.
// ctx.Scene is the scene being rendered.
func (t *testsModel) Initialize(ctx *Context) {

	// Show axis helper
	ah := graphic.NewAxisHelper(1.0)
	ctx.Scene.Add(ah)

	// Creates grid helper and saves its pointer in the test state
	t.grid = graphic.NewGridHelper(50, 1, &math32.Color{0.4, 0.4, 0.4})
	ctx.Scene.Add(t.grid)

	// Changes the camera position
	ctx.Camera.GetCamera().SetPosition(0, 4, 10)
}

// This method will be called at every frame
// You can animate your objects here.
func (t *testsModel) Render(ctx *Context) {

	// Rotate the grid, just for show.
	t.grid.AddRotationY(0.005)
}

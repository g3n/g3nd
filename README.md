# G3ND - G3N Game Engine Demo

G3ND is the demo for the [G3N](https://github.com/g3n/engine) 3D game engine.
It demonstrates and exercises the main features of the engine. Browsing and reading through the source code of the [demos](https://github.com/g3n/g3nd/tree/master/demos) is a great way to learn how to use the engine.
It's very easy to create a new demo as the main program takes care of initializing common objects and components.

<p align="center">
  <img style="float: right;" src="https://raw.githubusercontent.com/g3n/g3nd/master/data/images/g3nd_screenshots.png" alt="G3ND In Action"/>
</p>

# Dependencies

G3ND only depends on [G3N](https://github.com/g3n/engine) and so has the same dependencies as the engine itself.
Please verify that the [engine dependencies](https://github.com/g3n/engine#dependencies) are in place before installing.

# Installation

The following set of commands will download, compile, and install G3ND, the engine, and other Go packages on which the engine depends.
It will also generate the `g3nd` binary.

```
git clone https://github.com/g3n/g3nd
cd g3nd
go install
```

# Running

When G3ND is run without any command line parameters it shows the tree of
categorized available demos at the left of its window and an empty center area
to show the demo scene.
Click on a category in the tree to expand it and then select a demo to show.

At the upper right corner is located the `Control` folder, which when clicked
shows some controls which can change the parameters of the current demo.
To run G3ND at fullscreen press `Alt-F11` or start it using the `-fullscreen` command line flag.

To exit the program press ESC or close the window.

You can start G3ND to show a specific demo specifying the demo name (category plus "." plus name) in the command
line such as:

`>g3nd geometry.box`

The G3ND window shows the current FPS rate (frames per second) of your system and the maximum potential FPS rate.
The desired FPS rate can be adjusted using the command line parameters: `-swapinterval` and `-targetfps`.

# Creating a new demo/test

You can use the `tests/model.go` file as a template
for your tests. You can can change it directly or copy it to a
new file such as `tests/mytest.go` and
experiment with the engine. Your new test will appear under the
`|tests|` category with `mytest` name. The contents of the `tests/model.go`
file are shown below, documenting the common structure of all
demo programs:

```Go
// This is a simple model for your tests
package tests

import (
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/app"
	"time"
)

// Sets the category and name of your test in the demos.Map
// The category name choosen here starts with a "|" so it shows as the
// last category in list. Change "model" to the name of your test.
func init() {
	app.DemoMap["|tests|.model"] = &testsModel{}
}

// This is your test object. You can store state here.
// By convention and to avoid conflict with other demo/tests name it
// using your test category and name.
type testsModel struct {
	grid *graphic.GridHelper // Pointer to a GridHelper created in 'Start'
}

// This method will be called once when the test is selected from the G3ND list.
// 'a' is a pointer to the G3ND application.
// It allows access to several methods such as a.Scene(), which returns the current scene,
// a.DemoPanel(), a.Camera(), a.Window() among others.
// You can build your scene adding your objects to the a.Scene()
func (t *testsModel) Start(a *app.App) {

	// Show axis helper
	ah := graphic.NewAxisHelper(1.0)
	a.Scene().Add(ah)

	// Creates a grid helper and saves its pointer in the test state
	t.grid = graphic.NewGridHelper(50, 1, &math32.Color{0.4, 0.4, 0.4})
	a.Scene().Add(t.grid)

	// Changes the camera position
	a.Camera().GetCamera().SetPosition(0, 4, 10)
	a.Camera().GetCamera().LookAt(&math32.Vector3{0, 0, 0})
}

// This method will be called at every frame
// You can animate your objects here.
func (t *testsModel) Update(a *app.App, deltaTime time.Duration) {

	// Rotate the grid, just for show.
	rps := float32(deltaTime.Seconds()) * 2 * math32.Pi
	t.grid.RotateY(rps * 0.05)
}

// Cleanup is called once at the end of the demo.
func (t *testsModel) Cleanup(a *app.App) {}
```

# Contributing

If you spot a bug or create a new interesting demo you are encouraged to
send pull requests.

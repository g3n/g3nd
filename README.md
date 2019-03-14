# G3ND - G3N Game Engine Demo Program

G3ND is a demo/test program for the [G3N](https://github.com/g3n/engine) 3D game engine.
It demonstrates the main features of the engine and contains some basic tests.
It can be used to learn the game engine API by examining the source code of the demo programs.
It's very easy to create a new demo as the main program takes care of a lot of initializations and housekeeping.

<p align="center">
  <img style="float: right;" src="https://raw.githubusercontent.com/g3n/g3nd/master/data/images/g3nd_screenshots.png" alt="G3ND In Action"/>
</p>

# Dependencies for installation

G3ND only depends on the [G3N](https://github.com/g3n/engine) game engine and so has the same dependencies as the engine itself.
Please verify that the engine dependencies are in place before installing.

# Installation

The following command will download G3ND, the engine, and third party Go packages on which it depends.
It will also compile and install the packages and the `g3nd` binary.

## Install without Modules (Before Go 1.11)

Make sure your GOPATH is set correctly.

`go get -u github.com/g3n/g3nd`

## Install with Modules (Go 1.11 or higher)

GOPATH is not used with Go Modules.

    git clone https://github.com/g3n/g3nd
    cd g3nd
    go install 

Note: G3ND comes with a data directory with media files including images, textures, models and audio files.
Currently this directory has aproximately 50MB. The download and compilation may take some time.
To see what is going on you can alternatively supply the verbose flag:

`go get -u -v github.com/g3n/g3nd ;# without go modules`

`go get -u -v ;# with go modules, within g3nd directory`

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

```

# Contributing

If you spot a bug or create a new interesting demo you are encouraged to
send pull requests.



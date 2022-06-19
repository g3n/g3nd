package main

import (
	_ "github.com/g3n/g3nd/demos/animation"
	_ "github.com/g3n/g3nd/demos/audio"
	_ "github.com/g3n/g3nd/demos/experimental/physics"
	_ "github.com/g3n/g3nd/demos/geometry"
	_ "github.com/g3n/g3nd/demos/gui"
	_ "github.com/g3n/g3nd/demos/helper"
	_ "github.com/g3n/g3nd/demos/light"
	_ "github.com/g3n/g3nd/demos/loader"
	_ "github.com/g3n/g3nd/demos/material"
	_ "github.com/g3n/g3nd/demos/movement"
	_ "github.com/g3n/g3nd/demos/other"
	_ "github.com/g3n/g3nd/demos/shader"
	_ "github.com/g3n/g3nd/demos/tests"
	_ "github.com/g3n/g3nd/demos/texture"

	"github.com/g3n/g3nd/app"
)

func main() {

	// Create and run application
	app.Create().Run()
}

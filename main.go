package main

import (
	_ "github.com/g3n/g3nd/audio"
	_ "github.com/g3n/g3nd/geometry"
	_ "github.com/g3n/g3nd/gui"
	_ "github.com/g3n/g3nd/helper"
	_ "github.com/g3n/g3nd/light"
	_ "github.com/g3n/g3nd/loader"
	_ "github.com/g3n/g3nd/material"
	_ "github.com/g3n/g3nd/other"
	_ "github.com/g3n/g3nd/shader"
	_ "github.com/g3n/g3nd/skybox"
	_ "github.com/g3n/g3nd/texture"

	"github.com/g3n/g3nd/demos"
	"github.com/g3n/g3nd/g3nd"
)

func main() {

	app := g3nd.Create(demos.Map)
	if app != nil {
		app.Run()
	}
}

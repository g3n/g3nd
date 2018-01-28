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
	_ "github.com/g3n/g3nd/tests"
	_ "github.com/g3n/g3nd/texture"

	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/demos"
)

func main() {

	// Creates application and panics if error
	a := app.Create(demos.Map)
	if a == nil {
		return
	}
	// Runs application and panics if error
	err := a.Run()
	if err != nil {
		panic(err)
	}
}

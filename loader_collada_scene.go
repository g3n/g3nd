package main

import (
	"github.com/g3n/engine/loader/collada"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/light"
	"io"
)

func init() {
	TestMap["loader.collada_scene"] = &ColladaScene{}
}

type ColladaScene struct {
	normals *graphic.NormalsHelper
}

func (t *ColladaScene) Initialize(ctx *Context) {

	// Add directional top white light
	l1 := light.NewDirectional(&math32.Color{1, 1, 1}, 1.0)
	l1.SetPosition(0, 1, 0)
	ctx.Scene.Add(l1)

	// Add directional right white light
	l2 := light.NewDirectional(&math32.Color{1, 1, 1}, 1.0)
	l2.SetPosition(1, 0, 0)
	ctx.Scene.Add(l2)

	// Add directional front  white light
	l3 := light.NewDirectional(&math32.Color{1, 1, 1}, 1.0)
	l3.SetPosition(0, 1, 1)
	ctx.Scene.Add(l3)

	// Sets camera position
	ctx.Camera.GetCamera().SetPosition(0, 4, 10)

    // Adds axix helper
    ah := graphic.NewAxisHelper(1.5)
    ctx.Scene.Add(ah)

	dec, err := collada.Decode(ctx.DirData + "/collada/scene.dae")
	if err != nil && err != io.EOF {
		log.Fatal("%s", err)
	}
    dec.SetDirImages(ctx.DirData + "/images")

    // Loads collada scene
    s, err := dec.NewScene()
    if err != nil {
		log.Fatal("%s", err)
    }
    ctx.Scene.Add(s)
}

func (t *ColladaScene) Render(ctx *Context) {

}


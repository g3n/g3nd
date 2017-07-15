package main

import (
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/loader/gltf"
	"github.com/g3n/engine/math32"

	"github.com/davecgh/go-spew/spew"
)

func init() {
	TestMap["loader.gltf"] = &GltfLoader{}
}

type GltfLoader struct {
}

func (t *GltfLoader) Initialize(ctx *Context) {

	// Adds white directional front light
	l1 := light.NewDirectional(math32.NewColor(1, 1, 1), 1.0)
	l1.SetPosition(0, 0, 10)
	ctx.Scene.Add(l1)

	// Adds white directional top light
	l2 := light.NewDirectional(math32.NewColor(1, 1, 1), 1.0)
	l2.SetPosition(0, 10, 0)
	ctx.Scene.Add(l2)

	// Adds white directional right light
	l3 := light.NewDirectional(math32.NewColor(1, 1, 1), 1.0)
	l3.SetPosition(10, 0, 0)
	ctx.Scene.Add(l3)

	// Adds axis helper
	axis := graphic.NewAxisHelper(2)
	ctx.Scene.Add(axis)

	// Parses gltf file
	g, err := gltf.ParseJSON(ctx.DirData + "/gltf/triangle.gltf")
	if err != nil {
		log.Fatal(err.Error())
	}

	spew.Config.Indent = "   "
	spew.Dump(g.Meshes)
	spew.Dump(g.Accessors)

	n, err := g.NewNode(0)
	if err != nil {
		panic(err)
	}
	log.Error("node:%v", n)
	ctx.Scene.Add(n)
}

func (t *GltfLoader) Render(ctx *Context) {

}

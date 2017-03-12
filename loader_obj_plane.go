package main

import (
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/loader/obj"
	"github.com/g3n/engine/math32"
)

func init() {
	TestMap["loader.obj_plane"] = &ObjLoaderPlane{}
}

type ObjLoaderPlane struct {
}

func (t *ObjLoaderPlane) Initialize(ctx *Context) {

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

	// Decodes obj file and associated mtl file
	objfile := "spitfire.obj"
	dec, err := obj.Decode(ctx.DirData+"/obj/"+objfile, "")
	if err != nil {
		log.Fatal(err.Error())
	}

	// Creates a new node with all the objects in the decoded file and adds it to the scene
	group, err := dec.NewGroup()
	if err != nil {
		log.Fatal(err.Error())
	}
	ctx.Scene.Add(group)

	// Sets camera position
	ctx.Camera.GetCamera().SetPosition(-4, 3, -6)
}

func (t *ObjLoaderPlane) Render(ctx *Context) {

}

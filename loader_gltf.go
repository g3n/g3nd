package main

import (
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/loader/gltf"
	"github.com/g3n/engine/math32"
)

func init() {
	TestMap["loader.gltf"] = &GltfLoader{}
}

type GltfLoader struct {
	prevLoaded core.INode
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

	// Adds dropdown to select model to show
	dd := t.fileDropdown(ctx.DirData + "/gltf")
	dd.SetPosition(10, 10)
	dd.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		log.Debug("%v", dd.Selected().Text())
		t.loadModel(ctx, dd.Selected().Text())
	})
	ctx.Gui.Add(dd)
}

func (t *GltfLoader) Render(ctx *Context) {

}

func (t *GltfLoader) loadModel(ctx *Context, fname string) {

	// Remove previous model from the scene
	if t.prevLoaded != nil {
		ctx.Scene.Remove(t.prevLoaded)
		t.prevLoaded.Dispose()
		t.prevLoaded = nil
	}

	// Parses gltf file
	g, err := gltf.ParseJSON(ctx.DirData + "/gltf/" + fname)
	if err != nil {
		log.Fatal(err.Error())
	}

	spew.Config.Indent = "   "
	spew.Dump(g.Meshes)
	spew.Dump(g.Accessors)

	// Get node
	n, err := g.NewNode(0)
	if err != nil {
		panic(err)
	}
	log.Error("node:%v", n)
	ctx.Scene.Add(n)
	t.prevLoaded = n
}

func (t *GltfLoader) fileDropdown(dir string) *gui.DropDown {

	// Open dir
	f, err := os.Open(dir)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read files from directory
	files, err := f.Readdir(0)
	if err != nil {
		panic(err)
	}
	models := make([]string, 0)
	for _, fi := range files {
		if fi.IsDir() {
			continue
		}
		if strings.HasSuffix(fi.Name(), ".gltf") {
			models = append(models, fi.Name())
		}
	}

	// Creates DropDown
	dd := gui.NewDropDown(200, gui.NewImageLabel("Select Model"))
	for _, fname := range models {
		dd.Add(gui.NewImageLabel(fname))
	}
	return dd
}

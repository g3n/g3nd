package loader

import (
	"path/filepath"

	"github.com/g3n/engine/core"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/loader/obj"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/demos"
	"github.com/g3n/g3nd/g3nd"
	"github.com/g3n/g3nd/util"
)

func init() {
	demos.Map["loader.obj"] = &LoaderObj{}
}

type LoaderObj struct {
	prevLoaded core.INode
	selFile    *util.FileSelectButton
}

func (t *LoaderObj) Initialize(app *g3nd.App) {

	// Creates file selection button
	t.selFile = util.NewFileSelectButton(app.DirData()+"/obj", "Select File", 400, 300)
	t.selFile.SetPosition(10, 10)
	t.selFile.FS.SetFileFilters("*.obj")
	app.GuiPanel().Add(t.selFile)
	t.selFile.Subscribe("OnSelect", func(evname string, ev interface{}) {
		fpath := ev.(string)
		err := t.load(app, fpath)
		if err == nil {
			t.selFile.Label.SetText("File: " + filepath.Base(fpath))
			t.selFile.SetError("")
		} else {
			t.selFile.Label.SetText("Select File")
		}
	})

	// Adds white directional front light
	l1 := light.NewDirectional(&math32.Color{1, 1, 1}, 1.0)
	l1.SetPosition(0, 0, 10)
	app.Scene().Add(l1)

	// Adds white directional top light
	l2 := light.NewDirectional(&math32.Color{1, 1, 1}, 1.0)
	l2.SetPosition(0, 10, 0)
	app.Scene().Add(l2)

	// Adds white directional right light
	l3 := light.NewDirectional(&math32.Color{1, 1, 1}, 1.0)
	l3.SetPosition(10, 0, 0)
	app.Scene().Add(l3)

	// Adds axis helper
	axis := graphic.NewAxisHelper(2)
	app.Scene().Add(axis)

	fpath := "obj/cubemultitex.obj"
	t.load(app, filepath.Join(app.DirData(), fpath))
	t.selFile.Label.SetText("File: " + filepath.Base(fpath))
}

func (t *LoaderObj) load(app *g3nd.App, path string) error {

	// Remove previous model from the scene
	if t.prevLoaded != nil {
		app.Scene().Remove(t.prevLoaded)
		t.prevLoaded.Dispose()
		t.prevLoaded = nil
	}

	// Decodes obj file and associated mtl file
	dec, err := obj.Decode(path, "")
	if err != nil {
		t.selFile.SetError(err.Error())
		return err
	}

	// Creates a new node with all the objects in the decoded file and adds it to the scene
	group, err := dec.NewGroup()
	if err != nil {
		t.selFile.SetError(err.Error())
		return err
	}
	app.Scene().Add(group)
	t.prevLoaded = group
	return nil
}

func (t *LoaderObj) Render(app *g3nd.App) {

}

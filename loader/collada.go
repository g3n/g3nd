package loader

import (
	"io"
	"path/filepath"

	"github.com/g3n/engine/core"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/loader/collada"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/demos"
	"github.com/g3n/g3nd/g3nd"
	"github.com/g3n/g3nd/util"
)

func init() {
	demos.Map["loader.collada"] = &LoaderCollada{}
}

type LoaderCollada struct {
	prevLoaded  core.INode
	selFile     *util.FileSelectButton
	animTargets map[string]*collada.AnimationTarget
}

func (t *LoaderCollada) Initialize(app *g3nd.App) {

	// Creates file selection button
	t.selFile = util.NewFileSelectButton(app.DirData()+"/collada", "Select File", 400, 300)
	t.selFile.SetPosition(10, 10)
	t.selFile.FS.SetFileFilters("*.dae")
	app.GuiPanel().Add(t.selFile)
	t.selFile.Subscribe("OnSelect", func(evname string, ev interface{}) {
		fpath := ev.(string)
		err := t.load(app, fpath)
		if err == nil {
			t.selFile.Label.SetText("File: " + filepath.Base(fpath))
			t.selFile.SetError("")
		} else {
			t.selFile.Label.SetText("Select File")
			t.selFile.SetError(err.Error())
		}
	})

	// Add directional top white light
	l1 := light.NewDirectional(&math32.Color{1, 1, 1}, 1.0)
	l1.SetPosition(0, 1, 0)
	app.Scene().Add(l1)

	// Add directional right white light
	l2 := light.NewDirectional(&math32.Color{1, 1, 1}, 1.0)
	l2.SetPosition(1, 0, 0)
	app.Scene().Add(l2)

	// Add directional front  white light
	l3 := light.NewDirectional(&math32.Color{1, 1, 1}, 1.0)
	l3.SetPosition(0, 1, 1)
	app.Scene().Add(l3)

	// Sets camera position
	app.Camera().GetCamera().SetPosition(0, 4, 10)

	// Adds axix helper
	ah := graphic.NewAxisHelper(1.5)
	app.Scene().Add(ah)

	// Loads default model
	fpath := filepath.Join(app.DirData(), "collada/scene.dae")
	t.load(app, fpath)
	t.selFile.Label.SetText("File: " + filepath.Base(fpath))
}

func (t *LoaderCollada) load(app *g3nd.App, path string) error {

	// Remove previous model from the scene
	if t.prevLoaded != nil {
		app.Scene().Remove(t.prevLoaded)
		t.prevLoaded.Dispose()
		t.prevLoaded = nil
	}

	// Decodes collada file
	dec, err := collada.Decode(path)
	if err != nil && err != io.EOF {
		t.selFile.SetError(err.Error())
		return err
	}
	dec.SetDirImages(app.DirData() + "/images")

	// Loads collada scene
	s, err := dec.NewScene()
	if err != nil {
		t.selFile.SetError(err.Error())
		return err
	}
	app.Scene().Add(s)
	t.prevLoaded = s

	// Checks for animations
	ats, err := dec.NewAnimationTargets(s)
	if err == nil {
		t.animTargets = ats
		for _, at := range ats {
			at.SetStart(-1.0)
			at.Reset()
			at.SetLoop(true)
		}
	}
	return nil
}

func (t *LoaderCollada) Render(app *g3nd.App) {

	if t.animTargets != nil {
		dt := app.FrameDeltaSeconds()
		for _, at := range t.animTargets {
			at.Update(dt)
		}
	}
}

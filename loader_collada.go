package main

import (
	"io"
	"path/filepath"

	"github.com/g3n/engine/core"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/loader/collada"
	"github.com/g3n/engine/math32"
)

func init() {
	TestMap["loader.collada"] = &LoaderCollada{}
}

type LoaderCollada struct {
	prevLoaded  core.INode
	selFile     *FileSelectButton
	animTargets map[string]*collada.AnimationTarget
}

func (t *LoaderCollada) Initialize(ctx *Context) {

	// Creates file selection button
	t.selFile = NewFileSelectButton(ctx.DirData+"/collada", "Select File", 400, 300)
	t.selFile.SetPosition(10, 10)
	t.selFile.FS.SetFileFilters("*.dae")
	ctx.Gui.Add(t.selFile)
	t.selFile.Subscribe("OnSelect", func(evname string, ev interface{}) {
		fpath := ev.(string)
		err := t.load(ctx, fpath)
		if err == nil {
			t.selFile.Label.SetText("File: " + filepath.Base(fpath))
			t.selFile.SetError("")
		} else {
			t.selFile.Label.SetText("Select File")
		}
	})

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

	// Loads default model
	fpath := filepath.Join(ctx.DirData, "collada/scene.dae")
	t.load(ctx, fpath)
	t.selFile.Label.SetText("File: " + filepath.Base(fpath))
}

func (t *LoaderCollada) load(ctx *Context, path string) error {

	// Remove previous model from the scene
	if t.prevLoaded != nil {
		ctx.Scene.Remove(t.prevLoaded)
		t.prevLoaded.Dispose()
		t.prevLoaded = nil
	}

	// Decodes collada file
	dec, err := collada.Decode(path)
	if err != nil && err != io.EOF {
		t.selFile.SetError(err.Error())
		return err
	}
	dec.SetDirImages(ctx.DirData + "/images")

	// Loads collada scene
	s, err := dec.NewScene()
	if err != nil {
		t.selFile.SetError(err.Error())
		return err
	}
	ctx.Scene.Add(s)
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

func (t *LoaderCollada) Render(ctx *Context) {

	if t.animTargets != nil {
		dt := float32(ctx.TimeDelta.Seconds())
		for _, at := range t.animTargets {
			at.Update(dt)
		}
	}
}

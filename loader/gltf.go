package loader

import (
	"fmt"
	"path/filepath"
	"github.com/davecgh/go-spew/spew"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/graphic"


	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/loader/gltf"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/demos"
	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/util"
	"github.com/g3n/engine/animation"
)

func init() {
	demos.Map["loader.gltf"] = &GltfLoader{}
}

type GltfLoader struct {
	prevLoaded core.INode
	selFile    *util.FileSelectButton
	anims      []*animation.Animation
}

func (t *GltfLoader) Initialize(a *app.App) {

	// Creates file selection button
	t.selFile = util.NewFileSelectButton(a.DirData()+"/gltf", "Select File", 400, 300)
	t.selFile.SetPosition(10, 10)
	t.selFile.FS.SetFileFilters("*.gltf", "*.glb")
	a.GuiPanel().Add(t.selFile)
	t.selFile.Subscribe("OnSelect", func(evname string, ev interface{}) {
		fpath := ev.(string)
		err := t.loadScene(a, fpath)
		if err == nil {
			t.selFile.Label.SetText("File: " + filepath.Base(fpath))
			t.selFile.SetError("")
		} else {
			t.selFile.Label.SetText("Select File")
			t.selFile.SetError(err.Error())
		}
	})

	// Adds white directional front light
	l1 := light.NewDirectional(math32.NewColor("white"), 1.0)
	l1.SetPosition(0, 0, 10)
	a.Scene().Add(l1)

	// Adds white directional top light
	l2 := light.NewDirectional(math32.NewColor("white"), 1.0)
	l2.SetPosition(0, 10, 0)
	a.Scene().Add(l2)

	// Adds white directional right light
	l3 := light.NewDirectional(math32.NewColor("white"), 1.0)
	l3.SetPosition(10, 0, 0)
	a.Scene().Add(l3)

	// Adds axis helper
	axis := graphic.NewAxisHelper(2)
	a.Scene().Add(axis)

	// Label for error message
	errLabel := gui.NewLabel("")
	errLabel.SetFontSize(18)
	a.Gui().Add(errLabel)

	fpath := "gltf/DamagedHelmet/glTF/DamagedHelmet.gltf"
	t.loadScene(a, filepath.Join(a.DirData(), fpath))
	t.selFile.Label.SetText("File: " + filepath.Base(fpath))

}

func (t *GltfLoader) Render(a *app.App) {

	for i, anim := range t.anims {
		a.Log().Error("Animation %v", i)
		anim.Update(a.FrameDeltaSeconds())
	}
}

func (t *GltfLoader) loadScene(a *app.App, fpath string) error {

	// TODO move camera or scale scene such that it's nicely framed
	// TODO do this for other loaders as well

	// Remove previous model from the scene
	if t.prevLoaded != nil {
		t.anims = t.anims[:]
		a.Scene().Remove(t.prevLoaded)
		t.prevLoaded.Dispose()
		t.prevLoaded = nil
	}

	// Checks file extension
	ext := filepath.Ext(fpath)
	var g *gltf.GLTF
	var err error

	// Parses file
	if ext == ".gltf" {
		g, err = gltf.ParseJSON(fpath)
	} else if ext == ".glb" {
		g, err = gltf.ParseBin(fpath)
	} else {
		return fmt.Errorf("Unrecognized file extension:%s", ext)
	}

	if err != nil {
		return err
	}

	spew.Config.Indent = "   "
	spew.Dump(g.Nodes)
	spew.Dump(g.Meshes)
	spew.Dump(g.Accessors)

	defaultSceneIdx := 0
	if g.Scene != nil {
		defaultSceneIdx = *g.Scene
	}

	// Create default scene
	n, err := g.NewScene(defaultSceneIdx)
	if err != nil {
		return err
	}

	// Create animations
	for i := range g.Animations {
		anim, _ := g.NewAnimation(i)
		anim.SetLoop(true)
		t.anims = append(t.anims, anim)
	}

	// Add normals helper
	//box := n.GetNode().Children()[0].GetNode().Children()[0]
	//normals := graphic.NewNormalsHelper(box.(graphic.IGraphic), 0.1, &math32.Color{0, 0, 1}, 1)
	//a.Scene().Add(normals)

	a.Scene().Add(n)
	t.prevLoaded = n
	return nil
}

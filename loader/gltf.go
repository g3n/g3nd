package loader

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

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
)

func init() {
	demos.Map["loader.gltf"] = &GltfLoader{}
}

type GltfLoader struct {
	prevLoaded core.INode
	selFile    *util.FileSelectButton
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

	//// Creates file selector
	//fs := NewFileSelect(400, 300)
	//fs.SetVisible(false)
	//err := fs.SetFileFilters("*.gltf", "*.glb")
	//if err != nil {
	//	panic(err)
	//}
	//// Loads model when OK is clicked
	//fs.Subscribe("OnOK", func(evname string, ev interface{}) {
	//	fpath := fs.Selected()
	//	if fpath == "" {
	//		fs.SetVisible(false)
	//		return
	//	}
	//	err := t.loadScene(a, fpath)
	//	if err != nil {
	//		errLabel.SetText("ERROR: " + err.Error())
	//	} else {
	//		errLabel.SetText("")
	//	}
	//	fs.SetVisible(false)
	//})
	//// Hides file select Cancel is clicked
	//fs.Subscribe("OnCancel", func(evname string, ev interface{}) {
	//	fs.SetVisible(false)
	//})
	//a.Gui().Add(fs)
	//
	//// Adds button to open file selector
	//b := gui.NewButton("Select File")
	//b.SetPosition(10, 10)
	//b.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
	//	fs.SetPath(ctx.DirData + "/gltf")
	//	fs.SetVisible(true)
	//})
	//fs.SetPosition(b.Width()+20, b.Position().Y)
	//a.Gui().Add(b)

	// Sets error label position
	//errLabel.SetPosition(b.Width()+20, b.Position().Y)
}

func (t *GltfLoader) Render(a *app.App) {

}

func (t *GltfLoader) loadScene(a *app.App, fpath string) error {

	// Remove previous model from the scene
	if t.prevLoaded != nil {
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
	//spew.Dump(g.Meshes)
	//spew.Dump(g.Accessors)

	// Get node
	n, err := g.NewScene(0)
	if err != nil {
		return err
	}

	// Add normals helper
	//box := n.GetNode().Children()[0]
	//normals := graphic.NewNormalsHelper(box.(graphic.IGraphic), 0.5, &math32.Color{0, 0, 1}, 1)
	//ctx.Scene.Add(normals)

	a.Scene().Add(n)
	t.prevLoaded = n
	return nil
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

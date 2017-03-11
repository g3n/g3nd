package main

import (
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/loader/collada"
	"github.com/g3n/engine/math32"
	"io"
	"os"
)

func init() {
	TestMap["loader.collada_anim"] = &ColladaAnim{}
}

type ColladaAnim struct {
	normals     *graphic.NormalsHelper
	animTargets map[string]*collada.AnimationTarget
}

func (t *ColladaAnim) Initialize(ctx *Context) {

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

	//dec, err := collada.Decode(ctx.DirData + "/collada/cubeanim1.dae")
	//dec, err := collada.Decode(ctx.DirData + "/collada/ballanim1.dae")
	dec, err := collada.Decode(ctx.DirData + "/collada/animation1.dae")
	if err != nil && err != io.EOF {
		log.Fatal("%s", err)
	}

	if false {
		dec.Dump(os.Stdout, 0)
	}

	// Loads collada scene
	dec.SetDirImages(ctx.DirData + "/images")
	s, err := dec.NewScene()
	if err != nil {
		log.Fatal("%s", err)
	}
	ctx.Scene.Add(s)

	ats, err := dec.NewAnimationTargets(s)
	if err != nil {
		log.Fatal("%s", err)
	}
	t.animTargets = ats
	for _, at := range ats {
		at.SetStart(-1.0)
		at.Reset()
		at.SetLoop(true)
	}
}

func (t *ColladaAnim) Render(ctx *Context) {

	dt := float32(ctx.TimeDelta.Seconds())
	for _, at := range t.animTargets {
		at.Update(dt)
	}
}

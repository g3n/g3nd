package material

import (
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/loader/obj"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/util"
	"time"
)

func init() {
	app.DemoMap["material.physical_helmet"] = &PhysicalHelmet{}
}

type PhysicalHelmet struct {
	p1    *util.PointLightMesh
	p2    *util.PointLightMesh
	s1    *util.SpotLightMesh
	s2    *util.SpotLightMesh
	count float32
}

// Start is called once at the start of the demo.
func (t *PhysicalHelmet) Start(a *app.App) {

	// Add directional red light from right
	l1 := light.NewDirectional(&math32.Color{1, 0, 0}, 1.0)
	l1.SetPosition(0.1, 0, 0)
	a.Scene().Add(l1)

	// Add directional green light from top
	l2 := light.NewDirectional(&math32.Color{0, 1, 0}, 1.0)
	l2.SetPosition(0, 0.1, 0)
	a.Scene().Add(l2)

	// Add directional blue light from front
	l3 := light.NewDirectional(&math32.Color{0, 0, 1}, 1.0)
	l3.SetPosition(0, 0, 0.1)
	a.Scene().Add(l3)

	t.p1 = util.NewPointLightMesh(&math32.Color{1, 1, 1})
	a.Scene().Add(t.p1)

	t.p2 = util.NewPointLightMesh(&math32.Color{1, 0, 0})
	a.Scene().Add(t.p2)

	t.s1 = util.NewSpotLightMesh(&math32.Color{0, 0, 1})
	//t.s1.SetPosition(0.5, 1, 0.5)
	//a.Scene().Add(t.s1)

	t.s2 = util.NewSpotLightMesh(&math32.Color{1, 0, 0})
	//t.s2.SetPosition(-1, 0, 0)
	//t.s2.SetRotationZ(math.Pi/2)
	//a.Scene().Add(t.s2)

	// DamagedHelmet

	// Decodes obj file and associated mtl file
	dec, err := obj.Decode(a.DirData()+"/obj/DamagedHelmet.obj", "")
	if err != nil {
		panic(err)
	}

	// Creates a new node with all the objects in the decoded file and adds it to the scene
	geom, err := dec.NewGeometry(&dec.Objects[0])
	if err != nil {
		panic(err)
	}

	// Helper function to load texture and handle errors
	newTexture := func(path string) *texture.Texture2D {
		tex, err := texture.NewTexture2DFromImage(path)
		if err != nil {
			a.Log().Fatal("Error loading texture: %s", err)
		}
		tex.SetWrapS(gls.MIRRORED_REPEAT)
		tex.SetWrapT(gls.MIRRORED_REPEAT)
		return tex
	}

	pbrMat := material.NewPhysical()
	pbrMat.SetEmissiveFactor(math32.NewColor("white"))
	pbrMat.SetBaseColorMap(newTexture(a.DirData() + "/obj/DamagedHelmet_albedo.jpg"))
	pbrMat.SetMetallicRoughnessMap(newTexture(a.DirData() + "/obj/DamagedHelmet_metalRoughness.jpg"))
	pbrMat.SetNormalMap(newTexture(a.DirData() + "/obj/DamagedHelmet_normal.jpg"))
	pbrMat.SetEmissiveMap(newTexture(a.DirData() + "/obj/DamagedHelmet_emissive.jpg"))
	pbrMat.SetOcclusionMap(newTexture(a.DirData() + "/obj/DamagedHelmet_AO.jpg"))

	helmet := graphic.NewMesh(geom, pbrMat)
	a.Scene().Add(helmet)

}

// Update is called every frame.
func (t *PhysicalHelmet) Update(a *app.App, deltaTime time.Duration) {

	t.p1.SetPosition(float32(math32.Cos(t.count)), float32(math32.Sin(t.count)), 0)
	t.p2.SetPosition(0, 1.5*float32(math32.Sin(t.count)), 1.5*float32(math32.Cos(t.count)))

	t.s1.SetPosition(0, 1.5*float32(math32.Sin(t.count*1.5)), 1.5*float32(math32.Cos(t.count*1.5)))
	t.s2.SetPosition(1.5*float32(math32.Cos(t.count*1.5)), 1.5*float32(math32.Sin(t.count*1.5)), 0)

	t.count += 0.01
}

// Cleanup is called once at the end of the demo.
func (t *PhysicalHelmet) Cleanup(a *app.App) {}

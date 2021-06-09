package material

import (
	"math"
	"time"

	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/util"
)

func init() {
	app.DemoMap["material.physical_variations"] = &PhysicalVariations{}
}

type PhysicalVariations struct {
	p1    *util.PointLightMesh
	p2    *util.PointLightMesh
	s1    *util.SpotLightMesh
	s2    *util.SpotLightMesh
	d1    *light.Directional
	d2    *light.Directional
	d3    *light.Directional
	sX    *material.Standard
	sx    *material.Standard
	sY    *material.Standard
	sy    *material.Standard
	sZ    *material.Standard
	sz    *material.Standard
	count float64
}

// Start is called once at the start of the demo.
func (t *PhysicalVariations) Start(a *app.App) {

	a.AmbLight().SetIntensity(0.1)

	// Add directional red light from side
	t.d1 = light.NewDirectional(&math32.Color{1, 0, 0}, 1.0)
	t.d1.SetPosition(1, 0, 0)
	a.Scene().Add(t.d1)

	// Add directional green light from top
	t.d2 = light.NewDirectional(&math32.Color{0, 1, 0}, 1.0)
	t.d2.SetPosition(0, 1, 0)
	a.Scene().Add(t.d2)

	//Add directional blue light from front
	t.d3 = light.NewDirectional(&math32.Color{0, 0, 1}, 1.0)
	t.d3.SetPosition(0, 0, 1)
	a.Scene().Add(t.d3)

	t.p1 = util.NewPointLightMesh(&math32.Color{1, 1, 1})
	t.p1.Light.SetQuadraticDecay(0.1)
	t.p1.Light.SetLinearDecay(0.1)
	a.Scene().Add(t.p1)

	t.p2 = util.NewPointLightMesh(&math32.Color{1, 1, 1})
	t.p2.Light.SetQuadraticDecay(0.1)
	t.p2.Light.SetLinearDecay(0.1)
	a.Scene().Add(t.p2)

	t.s1 = util.NewSpotLightMesh(&math32.Color{0, 0, 1})
	t.s1.SetPosition(0, 3, 0)
	t.s1.SetVisible(false)
	a.Scene().Add(t.s1)

	t.s2 = util.NewSpotLightMesh(&math32.Color{1, 0, 0})
	t.s2.SetPosition(-3, 0, 0)
	t.s2.SetRotationZ(math.Pi / 2)
	t.s2.SetVisible(false)
	a.Scene().Add(t.s2)

	// Spheres
	sphereGeometry := geometry.NewSphere(0.4, 32, 16)

	sideNum := 6
	offset := (float32(sideNum)+1)/2.0 - 0.5
	for i := 1; i < sideNum; i += 1 {
		for j := 1; j < sideNum; j += 1 {
			for k := 1; k < sideNum; k += 1 {
				pbrMat := material.NewPhysical()
				pbrMat.SetMetallicFactor(float32(i) / float32(sideNum))
				pbrMat.SetRoughnessFactor(float32(j) / float32(sideNum))
				v := float32(k) / float32(sideNum)
				//pbrMat.SetEmissiveFactor(&math32.Color{0,0,v})
				pbrMat.SetBaseColorFactor(&math32.Color4{v, v, v, 1})
				sphere := graphic.NewMesh(sphereGeometry, pbrMat)
				sphere.SetPosition(float32(i)-offset, float32(j)-offset, float32(k)-offset)
				a.Scene().Add(sphere)
			}
		}
	}

	// Add labels/sprites
	createSprite := func(text string, pos *math32.Vector3) *material.Standard {
		font := gui.StyleDefault().Font
		font.SetPointSize(36)
		font.SetColor(&math32.Color4{1, 1, 1, 1})
		width, height := font.MeasureText(text)
		img := font.DrawText(text)
		tex := texture.NewTexture2DFromRGBA(img)

		plane_mat := material.NewStandard(math32.NewColor("white"))
		plane_mat.AddTexture(tex)
		plane_mat.SetTransparent(true)

		div := float32(100)
		sprite := graphic.NewSprite(float32(width)/div, float32(height)/div, plane_mat)
		sprite.SetPositionVec(pos)
		a.Scene().Add(sprite)

		return plane_mat
	}

	dist := float32(4)
	t.sX = createSprite("+ Metalness", &math32.Vector3{dist, 0, 0})
	t.sx = createSprite("- Metalness", &math32.Vector3{-dist, 0, 0})
	t.sY = createSprite("+ Roughness", &math32.Vector3{0, dist, 0})
	t.sy = createSprite("- Roughness", &math32.Vector3{0, -dist, 0})
	t.sZ = createSprite("+ Diffuse", &math32.Vector3{0, 0, dist})
	t.sz = createSprite("- Diffuse", &math32.Vector3{0, 0, -dist})

	// TODO adjust zoom level (need to implement OrbitControl.SetZoom())
	// a.Orbit().SetZoom()

	// Add controls
	if a.ControlFolder() == nil {
		return
	}
	gDirectional := a.ControlFolder().AddGroup("Directional lights")
	cb1 := gDirectional.AddCheckBox("Red").SetValue(t.d1.Visible())
	cb1.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		t.d1.SetVisible(!t.d1.Visible())
	})
	cb2 := gDirectional.AddCheckBox("Green").SetValue(t.d2.Visible())
	cb2.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		t.d2.SetVisible(!t.d2.Visible())
	})
	cb3 := gDirectional.AddCheckBox("Blue").SetValue(t.d3.Visible())
	cb3.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		t.d3.SetVisible(!t.d3.Visible())
	})
	gSpot := a.ControlFolder().AddGroup("Spot lights")
	cb4 := gSpot.AddCheckBox("Blue").SetValue(t.s1.Visible())
	cb4.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		t.s1.SetVisible(!t.s1.Visible())
	})
	cb5 := gSpot.AddCheckBox("Red").SetValue(t.s2.Visible())
	cb5.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		t.s2.SetVisible(!t.s2.Visible())
	})
}

// Update is called every frame.
func (t *PhysicalVariations) Update(a *app.App, deltaTime time.Duration) {

	// Rotate point lights around origin
	t.p1.SetPosition(4*float32(math.Cos(t.count)), 4*float32(math.Sin(t.count)), 0)
	t.p2.SetPosition(0, 4*float32(math.Cos(t.count*1.618)), 4*float32(math.Sin(t.count*1.618)))
	t.count += 0.02

	// Adjust transparency of sprites according to camera angle
	var camPos math32.Vector3
	a.Camera().WorldPosition(&camPos)
	camPos.Normalize()

	X := camPos.Dot(&math32.Vector3{1, 0, 0})
	x := camPos.Dot(&math32.Vector3{-1, 0, 0})
	t.sX.SetOpacity(1 - math32.Max(x, X))
	t.sx.SetOpacity(1 - math32.Max(x, X))

	Y := camPos.Dot(&math32.Vector3{0, 1, 0})
	y := camPos.Dot(&math32.Vector3{0, -1, 0})
	t.sY.SetOpacity(1 - math32.Max(y, Y))
	t.sy.SetOpacity(1 - math32.Max(y, Y))

	Z := camPos.Dot(&math32.Vector3{0, 0, 1})
	z := camPos.Dot(&math32.Vector3{0, 0, -1})
	t.sZ.SetOpacity(1 - math32.Max(z, Z))
	t.sz.SetOpacity(1 - math32.Max(z, Z))
}

// Cleanup is called once at the end of the demo.
func (t *PhysicalVariations) Cleanup(a *app.App) {}

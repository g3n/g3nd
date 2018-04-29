package other

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/demos"

	"math/rand"
	"github.com/g3n/engine/material"
)

func init() {
	demos.Map["other.performance"] = &Performance{}
}

type Performance struct {

}

func (t *Performance) Initialize(a *app.App) {

	torusGeometry := geometry.NewTorus(0.5, 0.2, 16, 16, 2*math32.Pi)

	halfSize := 20
	step := 2
	count := 0
	for i := -halfSize; i < (halfSize+1); i+=step {
		for j := -halfSize; j < (halfSize+1); j+=step {
			for k := -halfSize; k < (halfSize+1); k+=step {
				count += 1
				mat := material.NewStandard(&math32.Color{rand.Float32(), rand.Float32(), rand.Float32()})
				//mat.SetSpecularColor(math32.NewColor("white"))
				//mat.SetShininess(50)
				torus := graphic.NewMesh(torusGeometry, mat)
				torus.SetPosition(float32(i), float32(j), float32(k))
				torus.SetRotation(rand.Float32()*2*math32.Pi, rand.Float32()*2*math32.Pi, rand.Float32()*2*math32.Pi)
				//torus.Materials()[0].GetMaterial().GetMaterial().SetWireframe(true)
				a.Scene().Add(torus)
			}
		}
	}
	a.Log().Info("%v objects added to the scene!", count)

	//stepLight := 10
	//countLight := 0
	//for i := -halfSize; i < (halfSize+1); i+=stepLight {
	//	for j := -halfSize; j < (halfSize+1); j+=stepLight {
	//		for k := -halfSize; k < (halfSize+1); k+=stepLight {
	//			countLight += 1
	//			light := light.NewPoint(math32.NewColor("white"), 2.0)
	//			light.SetPosition(float32(i), float32(j), float32(k))
	//			light.SetLinearDecay(0.5)
	//			light.SetQuadraticDecay(0.5)
	//			a.Scene().Add(light)
	//		}
	//	}
	//}
	//a.Log().Info("%v lights added to the scene!", countLight)

}

func (t *Performance) Render(a *app.App) {

}

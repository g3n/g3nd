package other

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/app"
	"time"
)

func init() {
	app.DemoMap["other.curves"] = &Curves2{}
}

type Curves2 struct {
	points *graphic.Points
}

// Start is called once at the start of the demo.
func (t *Curves2) Start(a *app.App) {

	// Create geometry for continued curves
	geom1 := geometry.NewGeometry()
	positions := math32.NewArrayF32(0, 0)
	colors := math32.NewArrayF32(0, 0)
	controlPoints := []*math32.Vector3{}

	quadBezier := math32.NewBezierQuadratic(
		math32.NewVector3(-1,0,0),
		math32.NewVector3(-1,-1,0),
		math32.NewVector3(0,-1,0),
	30)
	controlPoints = append(controlPoints, 
		math32.NewVector3(-1,-1,0),
	)

	cubeBezier := math32.NewBezierCubic(
		math32.NewVector3(0,-1,0),
		math32.NewVector3(1,-1,0),
		math32.NewVector3(1,1,0),
		math32.NewVector3(0,1,0),
	30)
	controlPoints = append(controlPoints, 
		math32.NewVector3(1,-1,0),
		math32.NewVector3(1,1,0),
	)

	hermiteSpline := math32.NewHermiteSpline(
		math32.NewVector3(0,1,0),
		math32.NewVector3(0,2,0),
		math32.NewVector3(-1,0,0),
		math32.NewVector3(-2,0,0),
	30)
	controlPoints = append(controlPoints, 
		math32.NewVector3(0,2,0),
		math32.NewVector3(-2,0,0),
	)

	continuedCurve := quadBezier.Continue(cubeBezier).Continue(hermiteSpline)
	continuedPoints := continuedCurve.GetPoints() // 91 points

	for i := 0; i < len(continuedPoints); i++ {
		positions.AppendVector3(&continuedPoints[i])
		if i < 30 {
			colors.Append(1,0,0)
		} else if i <= 60 {
			colors.Append(0,1,0)
		} else {
			colors.Append(0,0,1)
		}
	}
	geom1.AddVBO(gls.NewVBO(positions).AddAttrib(gls.VertexPosition))
	geom1.AddVBO(gls.NewVBO(colors).AddAttrib(gls.VertexColor))
	mat1 := material.NewBasic()
	lines1 := graphic.NewLineStrip(geom1, mat1)
	a.Scene().Add(lines1)

	
	// Points from curve controls
	pointsGeom := geometry.NewGeometry()
	positions = math32.NewArrayF32(0, 0)
	for i := 0; i < len(controlPoints); i++ {
		positions.AppendVector3(controlPoints[i].Clone())
	}
	pointsGeom.AddVBO(gls.NewVBO(positions).AddAttrib(gls.VertexPosition))
	mat11 := material.NewPoint(&math32.Color{0, 0, 0})
	mat11.SetSize(50)
	points1 := graphic.NewPoints(pointsGeom, mat11)
	a.Scene().Add(points1)

	// CatmullRom Spline through control points
	geom2 := geometry.NewGeometry()
	positions = math32.NewArrayF32(0, 0)
	colors = math32.NewArrayF32(0, 0)

	catmullRom := math32.NewCatmullRomSpline(controlPoints, 30, true)
	catmullPoints := catmullRom.GetPoints()

	for i := 0; i < len(catmullPoints); i++ {
		positions.AppendVector3(&catmullPoints[i])
		if i % 3 == 0 {
			colors.Append(1,0,0)
			colors.Append(0,1,0)
			colors.Append(0,0,1)
		}
	}
	geom2.AddVBO(gls.NewVBO(positions).AddAttrib(gls.VertexPosition))
	geom2.AddVBO(gls.NewVBO(colors).AddAttrib(gls.VertexColor))
	mat2 := material.NewBasic()
	lines2 := graphic.NewLineStrip(geom2, mat2)
	a.Scene().Add(lines2)

}

// Update is called every frame.
func (t *Curves2) Update(a *app.App, deltaTime time.Duration) {}

// Cleanup is called once at the end of the demo.
func (t *Curves2) Cleanup(a *app.App) {}

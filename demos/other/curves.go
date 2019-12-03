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

type Curve struct {
	points []math32.Vector3
	length float32
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

	quadBezier := NewBezierQuadratic(
		math32.NewVector3(-1,0,0),
		math32.NewVector3(-1,-1,0),
		math32.NewVector3(0,-1,0),
	30)
	controlPoints = append(controlPoints, 
		math32.NewVector3(-1,-1,0),
	)

	cubeBezier := NewBezierCubic(
		math32.NewVector3(0,-1,0),
		math32.NewVector3(1,-1,0),
		math32.NewVector3(1,1,0),
		math32.NewVector3(0,1,0),
	30)
	controlPoints = append(controlPoints, 
		math32.NewVector3(1,-1,0),
		math32.NewVector3(1,1,0),
	)

	hermiteSpline := NewHermiteSpline(
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

	catmullRom := NewCatmullRomSpline(controlPoints, 30, true)
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


//// CURVES.GO CODE ////
// TODO: import package from g3n engine after PR

func (c *Curve) GetPoints() []math32.Vector3 {
	return c.points
}

func (c *Curve) GetLength() float32 {
	return c.length
}

func (c *Curve) SetLength() {
	points := c.points
	l := float32(0.0)
	for i := 1; i < len(points); i++ {
		p0 := points[i].Clone()
		p1 := points[i - 1].Clone()
		l += (p0.Sub(p1)).Length()
	}
	c.length = l
}

// Continue combines two curves
// creates and returns a pointer to a new curve
// combined curves are unaffected
func (c *Curve) Continue(other *Curve) *Curve {
	last := c.points[len(c.points) - 1].Clone()
	first := other.points[0].Clone()

	var continued, otherpoints []math32.Vector3
	for i := 0; i < len(c.points); i++ {
		continued = append(continued, *c.points[i].Clone())
	}
	for i := 1; i < len(other.points); i++ {
		otherpoints = append(otherpoints, *other.points[i].Clone())
	}
	for i := 0; i < len(otherpoints); i++ {
		continued = append(continued, *otherpoints[i].Sub(first).Add(last))
	}
	newC := new(Curve)
	newC.points = continued
	newC.SetLength()
	return newC
}

// NewBezierQuadratic creates and returns a pointer to a new curve
// Uses Vector3 pointers origin, control, and destination to calculate with
// int npoints as the desired number of points along the curve
func NewBezierQuadratic(origin, control, destination *math32.Vector3, npoints int) *Curve {
	c := new(Curve)

	if npoints <= 2 {
		npoints = 3
	}
	var equation func(float32, float32, float32, float32) float32
	equation = func(t, v0, v1, v2 float32) float32 {
		a0 := 1.0 - t
		result := a0 * a0 * v0 + 2.0 * t * a0 * v1 + t * t * v2
		return result
	}
	var bezier []math32.Vector3

	for i := 0; i <= npoints; i++ {
		t := float32(i) / float32(npoints)
		x := equation(t, origin.X, control.X, destination.X)
		y := equation(t, origin.Y, control.Y, destination.Y)
		z := equation(t, origin.Z, control.Z, destination.Z)
		vect := math32.NewVector3(x, y, z)
		bezier = append(bezier, *vect)
	}

	c.points = bezier
	c.SetLength()
	return c
}

// NewBezierCubic creates and returns a pointer to a new curve
// Uses Vector3 pointers origin, control1, control2, and destination to calculate with
// int npoints as the desired number of points along the curve
func NewBezierCubic(origin, control1, control2, destination *math32.Vector3, npoints int) *Curve {
	c := new(Curve)

	if npoints <= 3 {
		npoints = 4
	}
	
	var equation func(float32, float32, float32, float32, float32) float32
	equation = func(t, v0, v1, v2, v3 float32) float32 {
		a0 := 1.0 - t
		result := a0 * a0 * a0 * v0 + 3.0 * t * a0 * a0 * v1 + 3.0 * t * t * a0 * v2 + t * t * t * v3
		return result
	}
	var bezier []math32.Vector3

	for i := 0; i <= npoints; i++ {
		t := float32(i) / float32(npoints)
		x := equation(t, origin.X, control1.X, control2.X, destination.X)
		y := equation(t, origin.Y, control1.Y, control2.Y, destination.Y)
		z := equation(t, origin.Z, control1.Z, control2.Z, destination.Z)
		vect := math32.NewVector3(x, y, z)
		bezier = append(bezier, *vect)
	}

	c.points = bezier
	c.SetLength()
	return c
}

// NewHermiteSpline creates and returns a pointer to a new curve
// Uses Vector3 pointers origin, tangent1, destination, and tangent2 to calculate with
// int npoints as the desired number of points along the curve
func NewHermiteSpline(origin, tangent1, destination, tangent2 *math32.Vector3, npoints int) *Curve {
	c := new(Curve)

	var equation func(float32, *math32.Vector3, *math32.Vector3, *math32.Vector3, *math32.Vector3) *math32.Vector3
	equation = func(t float32, v0, tan0, v1, tan1 *math32.Vector3) *math32.Vector3 {
		t2 := t * t
		t3 := t * t2
		p0 := (2.0 * t3) - (3.0 * t2) + 1.0
		p1 := (-2.0 * t3) + (3.0 * t2)
		p2 := t3 - (2.0 * t2) + t
		p3 := t3 - t2
		x := (v0.X * p0) + (v1.X * p1) + (tan0.X * p2) + (tan1.X * p3)
		y := (v0.Y * p0) + (v1.Y * p1) + (tan0.Y * p2) + (tan1.Y * p3)
		z := (v0.Z * p0) + (v1.Z * p1) + (tan0.Z * p2) + (tan1.Z * p3)
		return math32.NewVector3(x, y, z)
	}

	step := float32(1.0) / float32(npoints)
	var hermite []math32.Vector3
	for i := 0; i <= npoints; i++ {
		vect := equation(float32(i) * step, origin, tangent1, destination, tangent2)
		hermite = append(hermite, *vect)
	}
	c.points = hermite
	c.SetLength()
	return c
}

// NewCatmullRomSpline creates and returns a pointer to a new curve
// Uses array of Vector3 pointers with int npoints as the desired number of points between supplied points
// Use Boolean closed with true to close the start and end points
func NewCatmullRomSpline(points []*math32.Vector3, npoints int, closed bool) *Curve {
	c := new(Curve)

	var equation func(float32, *math32.Vector3, *math32.Vector3, *math32.Vector3, *math32.Vector3) *math32.Vector3
	equation = func(t float32, v0, v1, v2, v3 *math32.Vector3) *math32.Vector3 {
		t2 := t * t;
		t3 := t * t2;
		x := 0.5 * ((((2.0 * v1.X) + ((-v0.X + v2.X) * t)) +
			(((((2.0 * v0.X) - (5.0 * v1.X)) + (4.0 * v2.X)) - v3.X) * t2)) +
			((((-v0.X + (3.0 * v1.X)) - (3.0 * v2.X)) + v3.X) * t3));
		y := 0.5 * ((((2.0 * v1.Y) + ((-v0.Y + v2.Y) * t)) +
			(((((2.0 * v0.Y) - (5.0 * v1.Y)) + (4.0 * v2.Y)) - v3.Y) * t2)) +
			((((-v0.Y + (3.0 * v1.Y)) - (3.0 * v2.Y)) + v3.Y) * t3));
		z := 0.5 * ((((2.0 * v1.Z) + ((-v0.Z + v2.Z) * t)) +
			(((((2.0 * v0.Z) - (5.0 * v1.Z)) + (4.0 * v2.Z)) - v3.Z) * t2)) +
			((((-v0.Z + (3.0 * v1.Z)) - (3.0 * v2.Z)) + v3.Z) * t3));
		return math32.NewVector3(x, y, z);
	}

	step := float32(1.0) / float32(npoints)
	var catmull []math32.Vector3
	var t float32
	if closed {
		count := len(points)
		for i := 0; i < count; i++ {
			t = 0.0
			for n := 0; n < npoints; n++ {
				vect := equation(t, points[i % count], points[(i + 1) % count], points[(i + 2) % count], points[(i + 3) % count])
				catmull = append(catmull, *vect)
				t += step
			}
		}
		catmull = append(catmull, catmull[0])
	} else {
		total := []*math32.Vector3{points[0].Clone()}
		total = append(total, points...)
		total = append(total, points[len(points) - 1].Clone())
		var i int
		for i = 0; i < len(total) - 3; i++ {
			t = 0
			for n := 0; n < npoints; n++ {
				vect := equation(t, total[i], total[i + 1], total[i + 2], total[i + 3])
				catmull = append(catmull, *vect)
				t += step
			}
		}
		i--
		vect := equation(t, total[i], total[i + 1], total[i + 2], total[i + 3])
		catmull = append(catmull, *vect)
	}
	c.points = catmull
	c.SetLength()
	return c
}

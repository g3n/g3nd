package main

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/text"
	"github.com/g3n/engine/texture"
)

func init() {
	TestMap["camera.circle_location"] = &CameraCircleLocation{}
}

type CameraCircleLocation struct {
	font          *text.Font
	boxLocationX  float32
	boxLocationY  float32
	boxLocationZ  float32
	rotationSpeed float32 //In degrees per second
	radian        float32
	radius        float32
}

func (t *CameraCircleLocation) Initialize(ctx *Context) {
	// Setup
	t.boxLocationX = 0
	t.boxLocationY = 0
	t.boxLocationZ = 0
	t.rotationSpeed = 40
	t.radius = 4
	t.radian = 0

	// Info
	label := gui.NewLabel("The box itself does not rotate, the camera rotates around it.")
	label.SetFontSize(16)
	label.SetPosition(10, 10)
	ctx.Gui.Add(label)

	// Axis helper
	ah := graphic.NewAxisHelper(1.0)
	ctx.Scene.Add(ah)

	// Create a box, code stolen from "other_text.go"
	// Creates Font
	fontfile := ctx.DirData + "/fonts/FreeSans.ttf"
	font, err := text.NewFont(fontfile)
	if err != nil {
		log.Fatal(err.Error())
	}
	font.SetLineSpacing(1.0)
	font.SetSize(28)
	font.SetDPI(72)
	font.SetFgColor4(&math32.Color4{0, 0, 1, 1})
	font.SetBgColor4(&math32.Color4{1, 1, 0, 0.8})
	t.font = font

	// Creates Box
	geom1 := geometry.NewBox(1, 1, 1, 2, 2, 2)
	mesh1 := graphic.NewMesh(geom1, nil)
	// Creates box materials
	width := 128
	height := 128
	faceNames := []string{"Right", "Left", "Top", "Bottom", "Front", "Back"}
	for idx, name := range faceNames {
		nwidth, nheight := font.MeasureText(name)
		fx := (width - nwidth) / 2
		fy := (height - nheight) / 2
		canvas := text.NewCanvas(width, height, math32.NewColor4(1, 1, 1, 0.8))
		canvas.DrawText(fx, fy, name, font)
		tex := texture.NewTexture2DFromRGBA(canvas.RGBA)
		mat := material.NewStandard(math32.NewColor(1, 1, 1))
		mat.AddTexture(tex)
		mesh1.AddGroupMaterial(mat, idx)
	}
	mesh1.SetPosition(t.boxLocationX, t.boxLocationY, t.boxLocationZ)
	ctx.Scene.Add(mesh1)

	// Point the camera to the box
	ctx.Camera.GetCamera().LookAt(&math32.Vector3{t.boxLocationX, t.boxLocationY, t.boxLocationY})
}

func (t *CameraCircleLocation) Render(ctx *Context) {
	// Move the camera arround the Y axis
	t.radian += t.rotationSpeed * float32(ctx.TimeDelta.Seconds()) * math32.Pi / 180 // In real applications a method should be in place to prevent an overflow
	cameraX := t.boxLocationX + t.radius*math32.Cos(t.radian)
	cameraZ := t.boxLocationZ + t.radius*math32.Sin(t.radian)
	ctx.Camera.GetCamera().SetPosition(float32(cameraX), t.boxLocationY, float32(cameraZ))
}

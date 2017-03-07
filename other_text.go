package main

import (
	"fmt"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/texture"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/text"
)

type Text1 struct {
	planeTime *graphic.Mesh
    planeTex  *texture.Texture2D
	font      *text.Font
}

func init() {
	TestMap["other.text"] = &Text1{}
}

// Draw the text.
var newtonMsg = `Sprite2:

We are merely picking up pebbles on the beach
while the great ocean of truth
lays completely undiscovered before us.

Isaac Newton.
`

func (t *Text1) Initialize(ctx *Context) {

	l1 := light.NewDirectional(math32.NewColor(1, 1, 1), 1.0)
	l1.SetPosition(0, 0, 10)
	ctx.Scene.Add(l1)

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
	ctx.Scene.Add(mesh1)

	// Plane texture
	canvas := text.NewCanvas(300, 200, math32.NewColor4(0, 1, 0, 0.8))
	canvas.DrawText(0, 20, "Message1", font)
	canvas.DrawText(100, 50, "Other text", font)
	font.SetFgColor4(&math32.Color4{1, 0, 0, 1})
	canvas.DrawText(20, 100, "In Red", font)
	font.SetSize(48)
	font.SetFgColor4(&math32.Color4{0, 0, 0, 1})
	canvas.DrawText(0, 150, "BIGGER", font)
	// Plane
	tex1 := texture.NewTexture2DFromRGBA(canvas.RGBA)
	geom2 := geometry.NewPlane(3, 2, 1, 1)
	mat2 := material.NewStandard(math32.NewColor(1, 1, 1))
	mat2.AddTexture(tex1)
	mesh2 := graphic.NewMesh(geom2, mat2)
	mesh2.SetPosition(2, 2, -0.8)
	ctx.Scene.Add(mesh2)

	// Sprite1
    stext := "Sprite1\nSprite1\nSprite1"
	swidth, sheight := font.MeasureText(stext)
	canvas = text.NewCanvas(swidth, sheight, math32.NewColor4(0, 1, 1, 1))
	canvas.DrawText(0, 0, stext, font)
	tex3 := texture.NewTexture2DFromRGBA(canvas.RGBA)
	mat3 := material.NewStandard(math32.NewColor(1, 1, 1))
	mat3.AddTexture(tex3)
	aspect := float32(swidth) / float32(sheight)
	mesh3 := graphic.NewSprite(aspect, 1, mat3)
	mesh3.SetPosition(-1.5, 1.5, 0.1)
	ctx.Scene.Add(mesh3)

	// Sprite2
	font.SetSize(28)
	swidth, sheight = font.MeasureText(newtonMsg)
	mx := 10
	swidth += 2 * mx
	canvas = text.NewCanvas(swidth, sheight, math32.NewColor4(1, 1, 1, 1))
	canvas.DrawText(mx, 0, newtonMsg, font)
	tex4 := texture.NewTexture2DFromRGBA(canvas.RGBA)
	mat4 := material.NewStandard(math32.NewColor(1, 1, 1))
	mat4.AddTexture(tex4)
	aspect = float32(swidth) / float32(sheight)
	mesh4 := graphic.NewSprite(aspect, 1, mat4)
	mesh4.SetPosition(1.5, -1.5, 0.1)
	ctx.Scene.Add(mesh4)

	// Plane2
	canvas = text.NewCanvas(256, 64, math32.NewColor4(1, 1, 1, 1))
	canvas.DrawText(0, 0, "", font)
	t.planeTex = texture.NewTexture2DFromRGBA(canvas.RGBA)
	geom5 := geometry.NewPlane(2, 0.5, 1, 1)
	mat5 := material.NewStandard(math32.NewColor(0.5, 0.8, 1))
	mat5.SetSide(material.SideDouble)
	mat5.AddTexture(t.planeTex)
	mesh5 := graphic.NewMesh(geom5, mat5)
	mesh5.SetPosition(-2, -1, -0.5)
	t.planeTime = mesh5
	ctx.Scene.Add(mesh5)

	axis := graphic.NewAxisHelper(1)
	ctx.Scene.Add(axis)
}

func (t *Text1) Render(ctx *Context) {

	l1 := fmt.Sprintf("Time: %4.4f", ctx.Win.GetTime())
	// Creates updated canvas
	canvas := text.NewCanvas(256, 64, math32.NewColor4(1, 1, 1, 1))
	t.font.SetSize(30)
	canvas.DrawText(0, 0, l1, t.font)
	// Update material texture
	t.planeTex.SetFromRGBA(canvas.RGBA)
	t.planeTime.AddRotationY(0.01)
}


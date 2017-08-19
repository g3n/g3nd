package main

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
)

type ShaderGeometry struct {
	ctx    *Context
	box    *graphic.Mesh
	sphere *graphic.Mesh
}

func init() {
	TestMap["shader.geometry"] = &ShaderGeometry{}
}

func (t *ShaderGeometry) Initialize(ctx *Context) {

	// Adds directional front light
	dir1 := light.NewDirectional(math32.NewColor(1, 1, 1), 0.6)
	dir1.SetPosition(0, 0, 100)
	ctx.Scene.Add(dir1)

	// Add axis helper
	axis := graphic.NewAxisHelper(1)
	ctx.Scene.Add(axis)

	err := ctx.Renderer.AddShader("shaderGeometryDemo", sourceGeometryDemo)
	if err != nil {
		panic(err)
	}
	err = ctx.Renderer.AddProgram("progGeometryDemo", "shaderStandardVertex", "shaderStandardFrag")
	if err != nil {
		panic(err)
	}
	err = ctx.Renderer.SetProgramShader("progGeometryDemo", gls.GEOMETRY_SHADER, "shaderGeometryDemo")
	if err != nil {
		panic(err)
	}

	// Creates box
	boxGeom := geometry.NewBox(1, 1, 1, 2, 2, 2)
	boxMat := NewNormalsMaterial(math32.NewColor(0.5, 0, 0))
	t.box = graphic.NewMesh(boxGeom, boxMat)
	t.box.SetPosition(-2, 0, 0)
	ctx.Scene.Add(t.box)

	// Creates sphere
	sphereGeom := geometry.NewSphere(1, 16, 16, 0, math32.Pi*2, 0, math32.Pi)
	sphereMat := NewNormalsMaterial(math32.NewColor(0, 0, 0.5))
	t.sphere = graphic.NewMesh(sphereGeom, sphereMat)
	t.sphere.SetPosition(2, 0, 0)
	ctx.Scene.Add(t.sphere)
}

func (t *ShaderGeometry) Render(ctx *Context) {

	t.box.AddRotationY(0.01)
	t.sphere.AddRotationZ(0.01)
}

//
// Normals Custom material
//

type NormalsMaterial struct {
	material.Standard // Embedded standard material
	vnormalColor      gls.Uniform3f
}

func NewNormalsMaterial(color *math32.Color) *NormalsMaterial {

	m := new(NormalsMaterial)
	m.Standard.Init("progGeometryDemo", color)

	// Creates uniforms
	m.vnormalColor.Init("VnormalColor")

	// Set initial values
	m.vnormalColor.SetColor(color)
	return m
}

func (m *NormalsMaterial) RenderSetup(gl *gls.GLS) {

	m.Standard.RenderSetup(gl)
	m.vnormalColor.Transfer(gl)
}

//
// Geometry Shader
//
const sourceGeometryDemo = `
#version {{.Version}}

layout (triangles) in;
layout (triangle_strip, max_vertices = 3) out;

// Outputs for the fragment shader.
out vec3 ColorFrontAmbdiff;
out vec3 ColorFrontSpec;
out vec3 ColorBackAmbdiff;
out vec3 ColorBackSpec;
out vec2 FragTexcoord;

void main(void) {

	gl_Position = gl_in[0].gl_Position;
    EmitVertex();

	gl_Position = gl_in[1].gl_Position;
    EmitVertex();

	gl_Position = gl_in[2].gl_Position;
    EmitVertex();

	EndPrimitive();
}

`

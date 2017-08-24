package main

import (
	"time"

	"github.com/g3n/engine/core"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
)

type OtherPemitter1 struct {
	pe1 *ParticleEmitter
	pe2 *ParticleEmitter
}

func init() {
	TestMap["other.pemitter1"] = &OtherPemitter1{}
}

func (t *OtherPemitter1) Initialize(ctx *Context) {

	// Sets background color
	ctx.GS.ClearColor(0, 0, 0, 1.0)

	// Add axis helper
	axis := graphic.NewAxisHelper(2)
	ctx.Scene.Add(axis)

	// Registers shaders and program
	err := ctx.Renderer.AddShader("shaderPEVertex", shaderPEVertex)
	if err != nil {
		panic(err)
	}
	err = ctx.Renderer.AddShader("shaderPEFrag", shaderPEFrag)
	if err != nil {
		panic(err)
	}
	err = ctx.Renderer.AddProgram("shaderPE", "shaderPEVertex", "shaderPEFrag")
	if err != nil {
		panic(err)
	}

	// Adds point particle emitter
	t.pe1 = NewParticleEmitter(2000)
	t.pe1.SetPosition(0, 0, 0)
	ctx.Scene.Add(t.pe1)

	t.pe2 = NewParticleEmitter(1000)
	t.pe2.SetPosition(1, 0, 0)
	t.pe2.AddRotationZ(-math32.Pi / 2)
	ctx.Scene.Add(t.pe2)

}

func (t *OtherPemitter1) Render(ctx *Context) {

	t.pe1.Update()
	t.pe2.Update()
}

//
//
// Implementation
//
//
type ParticleEmitter struct {
	graphic.Graphic                     // Embedded graphic
	mvpm            gls.UniformMatrix4f // Model view projection matrix uniform
	npoints         int
	start           time.Time
	mat             *ParticleEmitterMaterial
}

// NewParticleEmitter creates and returns a particle emitter object with the specified
// number of particles.
func NewParticleEmitter(npoints int) *ParticleEmitter {

	e := new(ParticleEmitter)
	e.npoints = npoints
	e.start = time.Now()

	// Creates geometry with points
	geom := geometry.NewGeometry()
	positions := math32.NewArrayF32(npoints*3, npoints*3)
	geom.AddVBO(gls.NewVBO().AddAttrib("VertexPosition", 3).SetBuffer(positions))
	e.Graphic.Init(geom, gls.POINTS)

	// Creates and add material
	e.mat = NewParticleEmitterMaterial()
	e.AddMaterial(e, e.mat, 0, 0)

	e.mvpm.Init("MVP")
	return e
}

func (e *ParticleEmitter) Update() {

	d := time.Now().Sub(e.start).Seconds()
	e.mat.PTime.Set(float32(d))

}

// RenderSetup is called by the engine before rendering this graphic
func (e *ParticleEmitter) RenderSetup(gs *gls.GLS, rinfo *core.RenderInfo) {

	// Calculates model view projection matrix and updates uniform
	mw := e.MatrixWorld()
	var mvpm math32.Matrix4
	mvpm.MultiplyMatrices(&rinfo.ViewMatrix, &mw)
	mvpm.MultiplyMatrices(&rinfo.ProjMatrix, &mvpm)
	e.mvpm.SetMatrix4(&mvpm)
	e.mvpm.Transfer(gs)
}

func (e *ParticleEmitter) Render(gs *gls.GLS) {

}

//
//
// Material
//
//
type ParticleEmitterMaterial struct {
	material.Material
	PTime gls.Uniform1f
}

func NewParticleEmitterMaterial() *ParticleEmitterMaterial {

	m := new(ParticleEmitterMaterial)
	m.Material.Init()
	m.SetShader("shaderPE")

	// Creates uniforms
	m.PTime.Init("PTime")

	// Set uniform's initial values
	m.PTime.Set(0)
	return m
}

func (m *ParticleEmitterMaterial) RenderSetup(gs *gls.GLS) {

	m.Material.RenderSetup(gs)
	m.PTime.Transfer(gs)
}

func (m *ParticleEmitterMaterial) Dispose() {

}

//
// Vertex Shader
//
const shaderPEVertex = `
#version {{.Version}}

// Uniform inputs
uniform mat4 MVP;
uniform float PTime;

// Output to fragment shader
smooth out vec4 vSmoothColor;


const vec3 a = vec3(0, 2, 0);	// acceleration 
const float rate = 1/500.0;     // rate of emission of particles
const float life = 2;			// particle life

const float PI = 3.14159;
const float TWO_PI = 2*PI;
const vec3 RED = vec3(1,0,0);
const vec3 GREEN = vec3(0,1,0);
const vec3 YELLOW = vec3(1,1,0);

// pseudorandom number generator
float rand(vec2 co){
	return fract(sin(dot(co.xy ,vec2(12.9898,78.233))) * 43758.5453);
}

// pseudorandom direction on a sphere
vec3 uniformRandomDir(vec2 v, out vec2 r) {
	
	r.x = rand(v.xy);
	r.y = rand(v.yx);
	float theta = mix(0.0, PI / 6.0, r.x);
	float phi = mix(0.0, TWO_PI, r.y);
	return vec3(sin(theta) * cos(phi), cos(theta), sin(theta) * sin(phi));
}

void main() {
	
	vec3 pos = vec3(0);
	float t = gl_VertexID * rate;
	float alpha = 1; 
	if (PTime > t) {
		float dt = mod((PTime - t), life);
		vec2 xy = vec2(gl_VertexID, t);
		vec2 rdm = vec2(0);
		pos = ((uniformRandomDir(xy, rdm) + 0.5 * a * dt) * dt);
		alpha = 1.0 - (dt / life);
	}
	vSmoothColor = vec4(mix(RED, YELLOW, alpha), alpha);
    gl_PointSize = 3;
	gl_Position = MVP * vec4(pos, 1);

}

`

//
// Fragment Shader
//
const shaderPEFrag = `
#version {{.Version}}

smooth in vec4 vSmoothColor;
layout(location=0) out vec4 vFragColor;

void main() {
	vFragColor = vSmoothColor;
}
`

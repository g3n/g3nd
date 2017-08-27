package main

import (
	"time"
	"unsafe"

	"github.com/g3n/engine/core"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
)

type OtherPemitter2 struct {
	pe1 *ParticleEmitter2
}

func init() {
	TestMap["other.pemitter2"] = &OtherPemitter2{}
}

func (t *OtherPemitter2) Initialize(ctx *Context) {

	// Sets background color
	ctx.GS.ClearColor(0, 0, 0, 1.0)

	// Add axis helper
	axis := graphic.NewAxisHelper(2)
	ctx.Scene.Add(axis)

	// Registers shaders and program
	err := ctx.Renderer.AddShader("shaderPE2Vertex", shaderPE2Vertex)
	if err != nil {
		panic(err)
	}
	err = ctx.Renderer.AddShader("shaderPE2Frag", shaderPE2Frag)
	if err != nil {
		panic(err)
	}
	err = ctx.Renderer.AddProgram("shaderPE2", "shaderPE2Vertex", "shaderPE2Frag")
	if err != nil {
		panic(err)
	}
	err = ctx.Renderer.SetProgramFeedbacks("shaderPE2",
		[]string{"OutPosition", "OutVelocity", "OutSTime"}, gls.INTERLEAVED_ATTRIBS)
	if err != nil {
		panic(err)
	}

	// Adds point particle emitter
	t.pe1 = NewParticleEmitter2(1)
	t.pe1.SetPosition(0, 0, 0)
	ctx.Scene.Add(t.pe1)
}

func (t *OtherPemitter2) Render(ctx *Context) {

	t.pe1.Update()
}

//
//
// Implementation
//
//
type ParticleEmitter2 struct {
	graphic.Graphic                     // Embedded graphic
	mvpm            gls.UniformMatrix4f // Model view projection matrix uniform
	npoints         int
	start           time.Time
	mat             *ParticleEmitter2Material
	gs              *gls.GLS        // gls state when initialized
	vboData         *gls.VBO        // VBO with points data
	data            math32.ArrayF32 // points data array
	feedbackHandle  uint32          // transform feedback buffer handle
	feedback        math32.ArrayF32 // transform feedback buffer
}

// NewParticleEmitter creates and returns a particle emitter object with the specified
// number of particles.
func NewParticleEmitter2(npoints int) *ParticleEmitter2 {

	e := new(ParticleEmitter2)
	e.npoints = npoints
	e.start = time.Now()

	// Creates particle data buffer: position(3) + velocity(3) + stime(1)
	nfloats := npoints*3*2 + 1
	e.data = math32.NewArrayF32(nfloats, nfloats)
	e.data.Set(0, 0.1, 0.2, 0.3)
	e.data.Set(3, 0.4, 0.5, 0.6)
	e.data.Set(6, 0.7)

	// Creates feedback buffer with the same sizes as the particle data buffer
	e.feedback = math32.NewArrayF32(nfloats, nfloats)

	// Creates point geometry and adds VBO
	geom := geometry.NewGeometry()
	e.vboData = gls.NewVBO()
	e.vboData.AddAttrib("Position", 3)
	e.vboData.AddAttrib("Velocity", 3)
	e.vboData.AddAttrib("STime", 1)
	e.vboData.SetBuffer(e.data)
	geom.AddVBO(e.vboData)

	// Creates this graphic object
	e.Graphic.Init(geom, gls.POINTS)

	// Creates and add material
	e.mat = NewParticleEmitter2Material()
	e.AddMaterial(e, e.mat, 0, 0)

	// Initialize uniforms
	e.mvpm.Init("MVP")

	return e
}

func (e *ParticleEmitter2) Update() {

	if e.gs == nil {
		return
	}

	// Updates particles time
	d := time.Now().Sub(e.start).Seconds()
	e.mat.PTime.Set(float32(d))

	// Reads transform feedback buffer written by shader
	e.gs.BindBufferBase(gls.TRANSFORM_FEEDBACK_BUFFER, 0, e.feedbackHandle)
	size := len(e.feedback) * int(unsafe.Sizeof(float32(0)))
	e.gs.GetBufferSubData(gls.TRANSFORM_FEEDBACK_BUFFER, 0, uint32(size), unsafe.Pointer(&e.feedback[0]))

	// Sends the feedback buffer as new data to the shader
	e.vboData.SetBuffer(e.feedback)
	log.Debug("feedback: %v/%v/%v", e.feedbackHandle, len(e.feedback), e.feedback)

}

// RenderSetup is called by the engine before rendering this graphic
func (e *ParticleEmitter2) RenderSetup(gs *gls.GLS, rinfo *core.RenderInfo) {

	if e.gs == nil {
		// Create transform feedback buffer
		e.feedbackHandle = gs.GenBuffer()
		gs.BindBuffer(gls.ARRAY_BUFFER, e.feedbackHandle)
		gs.BufferData(gls.ARRAY_BUFFER, len(e.feedback)*int(unsafe.Sizeof(float32(0))), nil, gls.STATIC_READ)
		//gs.BufferData(gls.ARRAY_BUFFER, len(e.feedback)*int(unsafe.Sizeof(float32(0))), nil, gls.STREAM_DRAW)
		e.gs = gs
		log.Debug("Create tfb:%v", e.feedbackHandle)
	}

	// Calculates model view projection matrix and updates uniform
	mw := e.MatrixWorld()
	var mvpm math32.Matrix4
	mvpm.MultiplyMatrices(&rinfo.ViewMatrix, &mw)
	mvpm.MultiplyMatrices(&rinfo.ProjMatrix, &mvpm)
	e.mvpm.SetMatrix4(&mvpm)
	e.mvpm.Transfer(gs)

	// Begin transform feeback
	gs.BindBufferBase(gls.TRANSFORM_FEEDBACK_BUFFER, 0, e.feedbackHandle)
	gs.BeginTransformFeedback(gls.POINTS)
}

// RenderEnd is called after rendering this graphic
func (e *ParticleEmitter2) RenderEnd(gs *gls.GLS) {

	gs.EndTransformFeedback()
}

//
//
// Material
//
//
type ParticleEmitter2Material struct {
	material.Material
	PTime gls.Uniform1f
}

func NewParticleEmitter2Material() *ParticleEmitter2Material {

	m := new(ParticleEmitter2Material)
	m.Material.Init()
	m.SetShader("shaderPE2")

	// Creates uniforms
	m.PTime.Init("PTime")
	m.PTime.Set(0)
	return m
}

func (m *ParticleEmitter2Material) RenderSetup(gs *gls.GLS) {

	m.Material.RenderSetup(gs)
	m.PTime.Transfer(gs)
}

func (m *ParticleEmitter2Material) Dispose() {

}

//
// Vertex Shader
//
const shaderPE2Vertex = `
#version {{.Version}}

// Points attribute inputs
in vec3  Position;		// particle position
in vec3  Velocity;		// particle velocity
in float STime;			// particle start time

// Points feedback outputs
out vec3  OutPosition;	// updated particle position
out vec3  OutVelocity;	// updated particle velocity
out float OutSTime;		// copy of particle start time

// Uniform inputs
uniform mat4 MVP;		// model view projection matrix
uniform float PTime;	// particles current time

// Output to fragment shader
smooth out vec4 vSmoothColor;

void main() {

	vec3 pos = Position;
	vec3 vel = Velocity;

	if (PTime >= STime) {
		pos.x += 0.01;
		pos.y += 0.01;
		pos.z += 0.01;

		vel.x += 0.01;
		vel.y += 0.01;
		vel.z += 0.01;
	}

	OutPosition = pos;
	OutVelocity = vel;
	OutSTime = STime;
	
    gl_PointSize = 5;
	gl_Position = MVP * vec4(pos, 1);
	vSmoothColor = vec4(1,1,1,1);
}
`

//
// Fragment Shader
//
const shaderPE2Frag = `
#version {{.Version}}

smooth in vec4 vSmoothColor;
layout(location=0) out vec4 vFragColor;

void main() {
	vFragColor = vSmoothColor;
}
`

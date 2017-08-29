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
		[]string{"OutPosition", "OutVelocity", "OutPTime"}, gls.INTERLEAVED_ATTRIBS)
	if err != nil {
		panic(err)
	}

	// Adds point particle emitter
	t.pe1 = NewParticleEmitter2(10)
	t.pe1.SetPosition(0, 0, 0)
	ctx.Scene.Add(t.pe1)
}

func (t *OtherPemitter2) Render(ctx *Context) {

	//t.pe1.SetPositionX(t.pe1.Position().X + 0.002)
	t.pe1.Update()
}

//
//
// Implementation
//
//
type ParticleEmitter2 struct {
	graphic.Graphic                           // Embedded graphic
	mw              gls.UniformMatrix4f       // Model world matrix uniform
	vp              gls.UniformMatrix4f       // View projection matrix uniform
	mvpm            gls.UniformMatrix4f       // Model view projection matrix uniform
	npoints         int                       // number of particles
	start           time.Time                 // start simulation time
	mat             *ParticleEmitter2Material // particles material
	gs              *gls.GLS                  // gls state when initialized
	vboData         *gls.VBO                  // VBO with points data
	data            math32.ArrayF32           // points data array
	feedbackHandle  uint32                    // transform feedback buffer handle
	feedback        math32.ArrayF32           // transform feedback buffer
}

// NewParticleEmitter creates and returns a particle emitter object with the specified
// number of particles.
func NewParticleEmitter2(npoints int) *ParticleEmitter2 {

	e := new(ParticleEmitter2)
	e.npoints = npoints
	e.start = time.Now()

	// Creates particle data buffer: position(3) + velocity(3) + ptime(1)
	nfloats := npoints*3*2 + 1
	e.data = math32.NewArrayF32(nfloats, nfloats)
	//e.data.Set(0, 0.1, 0.2, 0.3)
	//e.data.Set(3, 0, 0, 0)
	//e.data.Set(6, 1.0)

	// Creates feedback buffer with the same sizes as the particle data buffer
	e.feedback = math32.NewArrayF32(nfloats, nfloats)

	// Creates point geometry and adds VBO
	geom := geometry.NewGeometry()
	e.vboData = gls.NewVBO()
	e.vboData.AddAttrib("Position", 3)
	e.vboData.AddAttrib("Velocity", 3)
	e.vboData.AddAttrib("PTime", 1)
	e.vboData.SetBuffer(e.data)
	geom.AddVBO(e.vboData)

	// Creates this graphic object
	e.Graphic.Init(geom, gls.POINTS)

	// Creates and add material
	e.mat = NewParticleEmitter2Material()
	e.AddMaterial(e, e.mat, 0, 0)

	// Initialize uniforms
	e.mw.Init("MW")
	e.vp.Init("VP")
	e.mvpm.Init("MVP")

	return e
}

func (e *ParticleEmitter2) Update() {

	if e.gs == nil {
		return
	}

	// Updates particles simulation time uniform
	d := time.Now().Sub(e.start).Seconds()
	e.mat.STime.Set(float32(d))

	// Reads transform feedback buffer written by shader
	e.gs.BindBufferBase(gls.TRANSFORM_FEEDBACK_BUFFER, 0, e.feedbackHandle)
	size := len(e.feedback) * int(unsafe.Sizeof(float32(0)))
	e.gs.GetBufferSubData(gls.TRANSFORM_FEEDBACK_BUFFER, 0, uint32(size), unsafe.Pointer(&e.feedback[0]))

	// Sends the feedback buffer as new data to the shader
	e.vboData.SetBuffer(e.feedback)
	//log.Debug("stime:%v feedback: %v/%v/%v", e.mat.STime.Get(), e.feedbackHandle, len(e.feedback), e.feedback)

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
	}

	// Get matrix world and updates uniform
	mw := e.MatrixWorld()
	e.mw.SetMatrix4(&mw)
	e.mw.Transfer(gs)

	// Calculates view projection matrix and updates uniform
	var vp math32.Matrix4
	vp.MultiplyMatrices(&rinfo.ViewMatrix, &rinfo.ProjMatrix)
	e.vp.SetMatrix4(&vp)
	e.vp.Transfer(gs)

	// Calculates model view projection matrix and updates uniform
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
	STime gls.Uniform1f // simulation time
	PLife gls.Uniform1f // particles life time
}

func NewParticleEmitter2Material() *ParticleEmitter2Material {

	m := new(ParticleEmitter2Material)
	m.Material.Init()
	m.SetShader("shaderPE2")

	// Simulation time uniform
	m.STime.Init("STime")
	m.STime.Set(0)

	// Particles life uniform
	m.PLife.Init("PLife")
	m.PLife.Set(1.0)
	return m
}

func (m *ParticleEmitter2Material) RenderSetup(gs *gls.GLS) {

	m.Material.RenderSetup(gs)
	m.STime.Transfer(gs)
	m.PLife.Transfer(gs)
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
in float PTime;			// particle start time

// Points transform feedback outputs
out vec3  OutPosition;	// updated particle position
out vec3  OutVelocity;	// updated particle velocity
out float OutPTime;		// copy of particle start time

// Uniform inputs
uniform mat4 MW;		// model world matrix
uniform mat4 VP;		// view projection matrix
uniform mat4 MVP;		// model view projection matrix
uniform float STime;	// simulation time
uniform float PLife;	// particles life time in seconds

// Output to fragment shader
smooth out vec4 vSmoothColor;

// Global constants
const float PI = 3.14159;
const float TWO_PI = 2*PI;

// Forward functions declarations
void initialize(inout vec3 pos, inout vec3 vel);
void update(inout vec3 pos, inout vec3 vel);
vec3 randomDir(vec2 v, float angle);
float rand(vec2 co);


void main() {

	vec3 pos = Position;
	vec3 vel = Velocity;

	// Copy current particle start time to feedback buffer
	OutPTime = PTime;

	// Initialize particle (only once)
	if (PTime == 0) {
		initialize(pos, vel);
	}

	// If simulation time greater than current particle start time, particle may be active
	if (STime >= PTime) {
		// Calculates current particle life time adding some random effect
		vec2 xy = vec2(STime, gl_VertexID);
		float life = STime - PTime + rand(xy) * 0.1;
		// If current particle life time is less the particle life time,
		// this particle is active. Updates its position and velocity
		if (PLife > life) {
			update(pos, vel);
		// Particle is not active any more. Prepare for recycle
		} else {
			initialize(pos, vel);
		}
	}

	// Copy current particle position and velocity to feedback buffer
	OutPosition = pos;
	OutVelocity = vel;
	
    gl_PointSize = 5;
	gl_Position = MVP * vec4(pos, 1);
	vSmoothColor = vec4(1,1,1,1);
}

// Initialize particle
// Sets the initial position in world coordinates
// Sets random start velocity and random start time
void initialize(inout vec3 pos, inout vec3 vel) {

	// Resets start position
	pos = vec3(0);

	// Generates initial random velocity
	vec2 xy = vec2(STime, gl_VertexID);
	vec3 dir = randomDir(xy, PI/8); 
 	vel = 0.03 * dir;

	// Generates initial random start time
	OutPTime = STime + rand(xy) * 1;
}

// Updates particle position and velocity
void update(inout vec3 pos, inout vec3 vel) {

	pos = pos + vel;
}

// Generates pseudorandom direction on a sphere
// vec2 v: to generate random values
vec3 randomDir(vec2 v, float angle) {

	vec2 r;
	r.x = rand(v.xy);
	r.y = rand(v.yx);
	float theta = mix(0.0, angle, r.x);
	float phi = mix(0.0, TWO_PI, r.y);
	return vec3(sin(theta) * cos(phi), cos(theta), sin(theta) * sin(phi));
}

// Generates pseudorandom number
// vec2 v: to generate random values
float rand(vec2 v){

	return fract(sin(dot(v.xy ,vec2(12.9898,78.233))) * 43758.5453);
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

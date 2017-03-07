package main

import (
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
)

type ShaderBricks struct {
	ctx     *Context
	plane1  *graphic.Mesh
    box1    *graphic.Mesh
	sphere1 *graphic.Mesh
}

func init() {
	TestMap["shader.bricks"] = &ShaderBricks{}
}

func (t *ShaderBricks) Initialize(ctx *Context) {

	// Adds directional front light
	dir1 := light.NewDirectional(math32.NewColor(1, 1, 1), 0.6)
	dir1.SetPosition(0, 0, 100)
	ctx.Scene.Add(dir1)

	// Add axis helper
	axis := graphic.NewAxisHelper(1)
	ctx.Scene.Add(axis)

	// Create custom shader
    err := ctx.Renderer.AddShader("shaderBricksVertex", shaderBricksVertex)
    if err != nil {
        panic(err)
    }
	err = ctx.Renderer.AddShader("shaderBricksFrag", shaderBricksFrag)
    if err != nil {
        panic(err)
    }
	err = ctx.Renderer.AddProgram("shaderBricks", "shaderBricksVertex", "shaderBricksFrag")
    if err != nil {
        panic(err)
    }

	// Creates plane 1
	geom1 := geometry.NewPlane(2, 2, 4, 4)
	mat1 := NewBricksMaterial(math32.NewColor(0.8, 0.2, 0.1))
	mat1.SetSide(material.SideDouble)
    mat1.SetShininess(10)
    mat1.SetSpecularColor(math32.NewColor(0,0,0))
	t.plane1 = graphic.NewMesh(geom1, mat1)
	t.plane1.SetPosition(-1.2, 1, 0)
	ctx.Scene.Add(t.plane1)

    // Creates box1
	geom2:= geometry.NewBox(2, 2, 1, 2, 2, 2)
	mat2 := NewBricksMaterial(math32.NewColor(0.2, 0.4, 0.8))
	t.box1 = graphic.NewMesh(geom2, mat2)
	t.box1.SetPosition(1.2, 1, 0)
	ctx.Scene.Add(t.box1)

	// Creates sphere 1
	geom3 := geometry.NewSphere(1, 16, 16, 0, math32.Pi*2, 0, math32.Pi)
	mat3 := NewBricksMaterial(math32.NewColor(0.5, 0.6, 0.7))
	t.sphere1 = graphic.NewMesh(geom3, mat3)
	t.sphere1.SetPosition(0, -1.2, 0)
	ctx.Scene.Add(t.sphere1)
}

func (t *ShaderBricks) Render(ctx *Context) {

    t.plane1.AddRotationY(-0.005)
    t.box1.AddRotationY(0.01)
    t.sphere1.AddRotationZ(0.01)
}

//
// Bricks Custom material
//

type BricksMaterial struct {
	material.Standard // Embedded standard material
	brickColor        gls.Uniform3f
	mortarColor       gls.Uniform3f
	brickSize         gls.Uniform2f
	brickPercent      gls.Uniform2f
}

func NewBricksMaterial(color *math32.Color) *BricksMaterial {

	m := new(BricksMaterial)
	m.Standard.Init("shaderBricks", color)

	// Creates uniforms
	m.brickColor.Init("BrickColor")
	m.mortarColor.Init("MortarColor")
	m.brickSize.Init("BrickSize")
	m.brickPercent.Init("BrickPercent")

	// Set initial values
	m.brickColor.SetColor(color)
	m.mortarColor.Set(0.2, 0.3, 0.2)
	m.brickSize.Set(0.5, 0.2)
	m.brickPercent.Set(0.8, 0.8)
	return m
}

func (m *BricksMaterial) RenderSetup(gl *gls.GLS) {

    m.Standard.RenderSetup(gl)
	m.brickColor.Transfer(gl)
	m.mortarColor.Transfer(gl)
	m.brickSize.Transfer(gl)
	m.brickPercent.Transfer(gl)
}


//
// Vertex Shader
//
const shaderBricksVertex = `
#version {{.Version}}

{{template "attributes" .}}

// Model uniforms
uniform mat4 ModelViewMatrix;
uniform mat3 NormalMatrix;
uniform mat4 MVP;

{{template "material" .}}

// Output variables for Fragment shader
out vec4 Position;
out vec3 Normal;
out vec3 CamDir;
out vec2 FragTexcoord;
out vec2 VPosition;

void main() {

    // Transform this vertex position to camera coordinates.
    Position = ModelViewMatrix * vec4(VertexPosition, 1.0);
    VPosition = VertexPosition.xy;

    // Transform this vertex normal to camera coordinates.
    Normal = normalize(NormalMatrix * VertexNormal);

    // Calculate the direction vector from the vertex to the camera
    // The camera is at 0,0,0
    CamDir = normalize(-Position.xyz);

    // Flips texture coordinate Y if requested.
    vec2 texcoord = VertexTexcoord;
    {{ if .MatTexturesMax }}
    if (MatTexFlipY[0] > 0) {
        texcoord.y = 1 - texcoord.y;
    }
    {{ end }}
    FragTexcoord = texcoord;

    gl_Position = MVP * vec4(VertexPosition, 1.0);
}
`

//
// Fragment Shader
//
const shaderBricksFrag = `
#version {{.Version}}

// Inputs from vertex shader
in vec4 Position;       // Vertex position in camera coordinates.
in vec3 Normal;         // Vertex normal in camera coordinates.
in vec3 CamDir;         // Direction from vertex to camera
in vec2 FragTexcoord;
in vec2 VPosition;      // Vertex position in model coordinates (xy)

{{template "lights" .}}
{{template "material" .}}
{{template "phong_model" .}}

// Uniforms for configure brick pattern
uniform vec3 BrickColor;
uniform vec3 MortarColor;
uniform vec2 BrickSize;
uniform vec2 BrickPercent;

// Final fragment color
out vec4 FragColor;

void main() {

    vec2 position = VPosition / BrickSize;
    if (fract(position.y * 0.5) > 0.5) {
        position.x += 0.5;
    }
    position = fract(position);

    vec2 useBrick = step(position, BrickPercent);
    vec3 color = mix(MortarColor, BrickColor, useBrick.x * useBrick.y);

    // Combine material with brick pattern colors
    vec4 matDiffuse = vec4(color, 1.0);
    vec4 matAmbient = vec4(MatAmbientColor, MatOpacity) * vec4(color, 1.0);

    // Inverts the fragment normal if not FrontFacing
    vec3 fragNormal = Normal;
    if (!gl_FrontFacing) {
        fragNormal = -fragNormal;
    }

    // Calculates the Ambient+Diffuse and Specular colors for this fragment using the Phong model.
    vec3 Ambdiff, Spec;
    phongModel(Position, fragNormal, CamDir, vec3(matAmbient), vec3(matDiffuse), Ambdiff, Spec);

    // Final fragment color
    FragColor = min(vec4(Ambdiff + Spec, matDiffuse.a), vec4(1.0));
}

`


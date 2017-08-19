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

	err := ctx.Renderer.AddShader("shaderGSDemoVertex", sourceGSDemoVertex)
	if err != nil {
		panic(err)
	}
	err = ctx.Renderer.AddShader("shaderGSDemoGeometry", sourceGSDemoGeometry)
	if err != nil {
		panic(err)
	}
	err = ctx.Renderer.AddShader("shaderGSDemoFrag", sourceGSDemoFrag)
	if err != nil {
		panic(err)
	}
	err = ctx.Renderer.AddProgram("progGeometryDemo", "shaderGSDemoVertex", "shaderGSDemoFrag")
	if err != nil {
		panic(err)
	}
	err = ctx.Renderer.SetProgramShader("progGeometryDemo", gls.GEOMETRY_SHADER, "shaderGSDemoGeometry")
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

const sourceGSDemoVertex = `
#version {{.Version}}

{{template "attributes" .}}

// Model uniforms
uniform mat4 ModelViewMatrix;
uniform mat3 NormalMatrix;
uniform mat4 MVP;

{{template "lights" .}}
{{template "material" .}}
{{template "phong_model" .}}


// Outputs for the fragment shader.
out vec3 ColorFrontAmbdiff;
out vec3 ColorFrontSpec;
out vec3 ColorBackAmbdiff;
out vec3 ColorBackSpec;
out vec2 FragTexcoord;

void main() {

    // Transform this vertex normal to camera coordinates.
    vec3 normal = normalize(NormalMatrix * VertexNormal);

    // Calculate this vertex position in camera coordinates
    vec4 position = ModelViewMatrix * vec4(VertexPosition, 1.0);

    // Calculate the direction vector from the vertex to the camera
    // The camera is at 0,0,0
    vec3 camDir = normalize(-position.xyz);

    // Calculates the vertex Ambient+Diffuse and Specular colors using the Phong model
    // for the front and back
    phongModel(position,  normal, camDir, MatAmbientColor, MatDiffuseColor, ColorFrontAmbdiff, ColorFrontSpec);
    phongModel(position, -normal, camDir, MatAmbientColor, MatDiffuseColor, ColorBackAmbdiff, ColorBackSpec);

    vec2 texcoord = VertexTexcoord;
    {{if .MatTexturesMax }}
    // Flips texture coordinate Y if requested.
    if (MatTexFlipY(0)) {
        texcoord.y = 1 - texcoord.y;
    }
    {{ end }}
    FragTexcoord = texcoord;

    gl_Position = MVP * vec4(VertexPosition, 1.0);
}
`

//
// Geometry Shader
//
const sourceGSDemoGeometry = `
#version {{.Version}}

layout (triangles) in;
layout (triangle_strip, max_vertices = 3) out;

// Inputs from Vertex shader
in vec3 ColorFrontAmbdiff[];
in vec3 ColorFrontSpec[];
in vec3 ColorBackAmbdiff[];
in vec3 ColorBackSpec[];
in vec2 FragTexcoord[];

// Outputs for the fragment shader.
out vec3 fColorFrontAmbdiff;
out vec3 fColorFrontSpec;
out vec3 fColorBackAmbdiff;
out vec3 fColorBackSpec;
out vec2 fFragTexcoord;

void main(void) {

	int n;
	for (n = 0; n < gl_in.length(); n++) {
		gl_Position = gl_in[n].gl_Position;
		fColorFrontAmbdiff = ColorFrontAmbdiff[n];
		fColorFrontSpec = ColorFrontSpec[n];
		fColorBackAmbdiff = ColorBackAmbdiff[n];
		fColorBackSpec = ColorBackSpec[n];
		EmitVertex();
	}
	EndPrimitive();
}

`

//
// Fragment Shader template
//
const sourceGSDemoFrag = `
#version {{.Version}}

{{template "material" .}}

// Inputs from Vertex shader
in vec3 fColorFrontAmbdiff;
in vec3 fColorFrontSpec;
in vec3 fColorBackAmbdiff;
in vec3 fColorBackSpec;
in vec2 fFragTexcoord;

// Output
out vec4 FragColor;


void main() {

    vec4 texCombined = vec4(1);

    // Combine all texture colors and opacity
    // Use Go templates to unroll the loop because non-const
    // array indexes are not allowed until GLSL 4.00.
    {{ range loop .MatTexturesMax }}
    if (MatTexVisible({{.}})) {
        vec4 texcolor = texture(MatTexture[{{.}}], fFragTexcoord * MatTexRepeat({{.}}) + MatTexOffset({{.}}));
        if ({{.}} == 0) {
            texCombined = texcolor;
        } else {
            texCombined = mix(texCombined, texcolor, texcolor.a);
        }
    }
    {{ end }}

    vec4 colorAmbDiff;
    vec4 colorSpec;
    if (gl_FrontFacing) {
        colorAmbDiff = vec4(fColorFrontAmbdiff, MatOpacity);
        colorSpec = vec4(fColorFrontSpec, 0);
    } else {
        colorAmbDiff = vec4(fColorBackAmbdiff, MatOpacity);
        colorSpec = vec4(fColorBackSpec, 0);
    }
    FragColor = min(colorAmbDiff * texCombined + colorSpec, vec4(1));
}

`

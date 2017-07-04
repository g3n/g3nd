package main

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
)

type Earth struct {
	ctx    *Context
	sphere *graphic.Mesh
}

func init() {
	TestMap["shader.earth"] = &Earth{}
}

func (t *Earth) Initialize(ctx *Context) {

	t.ctx = ctx
	ctx.GS.ClearColor(0, 0, 0, 1)

	ctx.AmbLight.SetIntensity(1)

	// Create Skybox
	skyboxData := graphic.SkyboxData{
		ctx.DirData + "/images/space/dark-s_", "jpg",
		[6]string{"px", "nx", "py", "ny", "pz", "nz"}}
	skybox, err := graphic.NewSkybox(skyboxData)
	if err != nil {
		panic(err)
	}
	ctx.Scene.Add(skybox)

	// Adds directional front light
	dir1 := light.NewDirectional(math32.NewColor(1, 1, 1), 0.9)
	dir1.SetPosition(0, 0, 100)
	ctx.Scene.Add(dir1)

	// Create day texture
	texDay, err := texture.NewTexture2DFromImage(ctx.DirData + "/images/earth_clouds_big.jpg")
	if err != nil {
		log.Fatal("Error loading texture: %s", err)
	}
	texDay.SetFlipY(false)

	// Create specular map texture
	texSpecular, err := texture.NewTexture2DFromImage(ctx.DirData + "/images/earth_spec_big.jpg")
	if err != nil {
		log.Fatal("Error loading texture: %s", err)
	}
	texSpecular.SetFlipY(false)

	// Create night texture
	texNight, err := texture.NewTexture2DFromImage(ctx.DirData + "/images/earth_night_big.jpg")
	if err != nil {
		log.Fatal("Error loading texture: %s", err)
	}
	texNight.SetFlipY(false)

	// Create bump map texture
	texBump, err := texture.NewTexture2DFromImage(ctx.DirData + "/images/earth_bump_big.jpg")
	if err != nil {
		log.Fatal("Error loading texture: %s", err)
	}
	texBump.SetFlipY(false)

	// Create custom shader
	err = t.ctx.Renderer.AddShader("shaderEarthVertex", shaderEarthVertex)
	if err != nil {
		panic(err)
	}
	err = t.ctx.Renderer.AddShader("shaderEarthFrag", shaderEarthFrag)
	if err != nil {
		panic(err)
	}
	err = t.ctx.Renderer.AddProgram("shaderEarth", "shaderEarthVertex", "shaderEarthFrag")
	if err != nil {
		panic(err)
	}

	// Create custom material using the custom shader
	matEarth := NewEarthMaterial(math32.NewColor(1, 1, 1))
	matEarth.SetShininess(20)
	//matEarth.SetSpecularColor(&math32.Color{0., 1, 1})
	//matEarth.SetColor(&math32.Color{0.8, 0.8, 0.8})

	// Create sphere
	geom := geometry.NewSphere(1, 32, 32, 0, math32.Pi*2, 0, math32.Pi)
	matEarth.AddTexture(texDay)
	matEarth.AddTexture(texSpecular)
	matEarth.AddTexture(texNight)
	t.sphere = graphic.NewMesh(geom, matEarth)
	ctx.Scene.Add(t.sphere)

	// Add axis helper
	axis := graphic.NewAxisHelper(5)
	ctx.Scene.Add(axis)

}

func (t *Earth) Render(ctx *Context) {

	t.sphere.AddRotationY(0.1 * float32(ctx.TimeDelta.Seconds()))
}

//
// Custom material
//

type EarthMaterial struct {
	material.Standard // Embedded standard material
}

// NewEarthMaterial creates and returns a pointer to a new earth material
func NewEarthMaterial(color *math32.Color) *EarthMaterial {

	pm := new(EarthMaterial)
	pm.Standard.Init("shaderEarth", color)
	return pm
}

//
// Vertex Shader
//
const shaderEarthVertex = `
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

out vec4 worldPosition;

void main() {

	worldPosition = vec4(VertexPosition, 1.0);

    // Transform this vertex position to camera coordinates.
    Position = ModelViewMatrix * vec4(VertexPosition, 1.0);

    // Transform this vertex normal to camera coordinates.
    Normal = normalize(NormalMatrix * VertexNormal);

    // Calculate the direction vector from the vertex to the camera
    // The camera is at 0,0,0
    CamDir = normalize(-Position.xyz);

    // Flips texture coordinate Y if requested.
    vec2 texcoord = VertexTexcoord;
    {{ if .MatTexturesMax }}
    if (MatTexFlipY(0)) {
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
const shaderEarthFrag = `
#version {{.Version}}

// Inputs from vertex shader
in vec4 Position;       // Vertex position in camera coordinates.
in vec3 Normal;         // Vertex normal in camera coordinates.
in vec3 CamDir;         // Direction from vertex to camera
in vec2 FragTexcoord;

in vec4 worldPosition;

{{template "lights" .}}
{{template "material" .}}
{{template "phong_model" .}}

// Final fragment color
out vec4 FragColor;

void logisticInterp(vec4 a, vec4 b, float f, out float r) {
	
}

void main() {

    vec4 texDay = texture(MatTexture[0], FragTexcoord * MatTexRepeat(0) + MatTexOffset(0));
    vec4 texSpecular = texture(MatTexture[1], FragTexcoord * MatTexRepeat(1) + MatTexOffset(1));
    vec4 texNight = texture(MatTexture[2], FragTexcoord * MatTexRepeat(2) + MatTexOffset(2));

    vec3 sunDirection = normalize(DirLightPosition(0));

    //vec4 texDayOrNight;// = texDay;

    // Inverts the fragment normal if not FrontFacing
    vec3 fragNormal = Normal;
    if (!gl_FrontFacing) {
        fragNormal = -fragNormal;
    }

    float dotNormal = dot(sunDirection, fragNormal);
    //if (dotNormal < 0) {
    //	texDayOrNight = texNight;
    //}

    vec4 texDayOrNight = mix(texNight, texDay, max(min((((dotNormal + 1.0)/2.0) - 0.45)*10.0, 1.0), 0.0)  );

    // Combine material with texture colors
    vec4 matDiffuse = vec4(MatDiffuseColor, MatOpacity) * texDayOrNight;
    vec4 matAmbient = vec4(MatAmbientColor, MatOpacity) * texDayOrNight;

    // Calculates the Ambient+Diffuse and Specular colors for this fragment using the Phong model.
    vec3 Ambdiff, Spec;
    phongModel(Position, fragNormal, CamDir, vec3(matAmbient), vec3(matDiffuse), Ambdiff, Spec);

    // Calculate specular mask
    Spec = vec3(texSpecular) * Spec;

    // Final fragment color
    FragColor = min(vec4(Ambdiff + Spec, matDiffuse.a), vec4(1.0));
}

`

package shader

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/demos"
)

type Earth struct {
	a      *app.App
	sphere *graphic.Mesh
}

func init() {
	demos.Map["shader.earth"] = &Earth{}
}

func (t *Earth) Initialize(a *app.App) {

	t.a = a
	a.AmbLight().SetIntensity(1)

	// Create Skybox
	skyboxData := graphic.SkyboxData{
		a.DirData() + "/images/space/dark-s_", "jpg",
		[6]string{"px", "nx", "py", "ny", "pz", "nz"}}
	skybox, err := graphic.NewSkybox(skyboxData)
	if err != nil {
		panic(err)
	}
	skybox.SetRenderOrder(-1) // The skybox should always be rendered first
	// For each skybox face sets the material to not use lights
	// and to have emissive color.
	sbmats := skybox.Materials()
	for i := 0; i < len(sbmats); i++ {
		sbmat := skybox.Materials()[i].GetMaterial().(*material.Standard)
		sbmat.SetUseLights(material.UseLightNone)
		sbmat.SetEmissiveColor(&math32.Color{1, 1, 1})
	}
	a.Scene().Add(skybox)

	// Adds directional front light
	dir1 := light.NewDirectional(&math32.Color{1, 1, 1}, 0.9)
	dir1.SetPosition(0, 0, 100)
	a.Scene().Add(dir1)

	// Create day texture
	texDay, err := texture.NewTexture2DFromImage(a.DirData() + "/images/earth_clouds_big.jpg")
	if err != nil {
		a.Log().Fatal("Error loading texture: %s", err)
	}
	texDay.SetFlipY(false)

	// Create specular map texture
	texSpecular, err := texture.NewTexture2DFromImage(a.DirData() + "/images/earth_spec_big.jpg")
	if err != nil {
		a.Log().Fatal("Error loading texture: %s", err)
	}
	texSpecular.SetFlipY(false)

	// Create night texture
	texNight, err := texture.NewTexture2DFromImage(a.DirData() + "/images/earth_night_big.jpg")
	if err != nil {
		a.Log().Fatal("Error loading texture: %s", err)
	}
	texNight.SetFlipY(false)

	// Create bump map texture
	texBump, err := texture.NewTexture2DFromImage(a.DirData() + "/images/earth_bump_big.jpg")
	if err != nil {
		a.Log().Fatal("Error loading texture: %s", err)
	}
	texBump.SetFlipY(false)

	// Create custom shader
	t.a.Renderer().AddShader("shaderEarthVertex", shaderEarthVertex)
	t.a.Renderer().AddShader("shaderEarthFrag", shaderEarthFrag)
	t.a.Renderer().AddProgram("shaderEarth", "shaderEarthVertex", "shaderEarthFrag")

	// Create custom material using the custom shader
	matEarth := NewEarthMaterial(&math32.Color{1, 1, 1})
	matEarth.SetShininess(20)
	//matEarth.SetSpecularColor(&math32.Color{0., 1, 1})
	//matEarth.SetColor(&math32.Color{0.8, 0.8, 0.8})

	// Create sphere
	geom := geometry.NewSphere(1, 32, 32, 0, math32.Pi*2, 0, math32.Pi)
	matEarth.AddTexture(texDay)
	matEarth.AddTexture(texSpecular)
	matEarth.AddTexture(texNight)
	t.sphere = graphic.NewMesh(geom, matEarth)
	a.Scene().Add(t.sphere)

	// Create sun sprite
	texSun, err := texture.NewTexture2DFromImage(a.DirData() + "/images/lensflare0_alpha.png")
	if err != nil {
		a.Log().Fatal("Error loading texture: %s", err)
	}
	sunMat := material.NewStandard(&math32.Color{1, 1, 1})
	sunMat.AddTexture(texSun)
	sunMat.SetTransparent(true)
	sun := graphic.NewSprite(10, 10, sunMat)
	sun.SetPositionZ(20)
	a.Scene().Add(sun)

	// Add axis helper
	axis := graphic.NewAxisHelper(5)
	a.Scene().Add(axis)

}

func (t *Earth) Render(a *app.App) {

	t.sphere.AddRotationY(0.1 * a.FrameDeltaSeconds())
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
#include <attributes>
#include <material>

// Model uniforms
uniform mat4 ModelViewMatrix;
uniform mat3 NormalMatrix;
uniform mat4 MVP;

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
	#if MAT_TEXTURES>0
    if (MatTexFlipY(0)) {
        texcoord.y = 1 - texcoord.y;
    }
	#endif
    FragTexcoord = texcoord;

    gl_Position = MVP * vec4(VertexPosition, 1.0);
}
`

//
// Fragment Shader
//
const shaderEarthFrag = `
#include <lights>
#include <material>
#include <phong_model>

// Inputs from vertex shader
in vec4 Position;       // Vertex position in camera coordinates.
in vec3 Normal;         // Vertex normal in camera coordinates.
in vec3 CamDir;         // Direction from vertex to camera
in vec2 FragTexcoord;

in vec4 worldPosition;

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

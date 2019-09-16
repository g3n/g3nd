package other

import (
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
	"github.com/g3n/engine/util"
	"github.com/g3n/g3nd/app"

	"time"
)

func init() {
	app.DemoMap["other.sprite_anim"] = &SpriteAnim{}
}

type SpriteAnim struct {
	anims []*texture.Animator
}

// Start is called once at the start of the demo.
func (t *SpriteAnim) Start(a *app.App) {

	a.AmbLight().SetIntensity(1)

	// Initialize list of animators
	t.anims = make([]*texture.Animator, 0)

	// Adds axis helper
	axis := util.NewAxisHelper(2)
	a.Scene().Add(axis)

	// Creates texture 1 and animator
	tex1, err := texture.NewTexture2DFromImage(a.DirData() + "/images/explosion7.png")
	if err != nil {
		a.Log().Fatal("Error loading texture: %s", err)
	}
	anim1 := texture.NewAnimator(tex1, 8, 8)
	anim1.SetDispTime(16666 * time.Microsecond)
	anim1.SetMaxCycles(4)
	t.anims = append(t.anims, anim1)

	mat1 := material.NewStandard(&math32.Color{1, 1, 1})
	mat1.AddTexture(tex1)
	mat1.SetOpacity(1)
	mat1.SetTransparent(true)
	s1 := graphic.NewSprite(2, 2, mat1)
	s1.SetPosition(-2, 2, 0)
	a.Scene().Add(s1)

	// Creates texture 2 and animator
	tex2, err := texture.NewTexture2DFromImage(a.DirData() + "/images/smoke30.png")
	if err != nil {
		a.Log().Fatal("Error loading texture: %s", err)
	}
	anim2 := texture.NewAnimator(tex2, 6, 5)
	anim2.SetDispTime(2 * 16666 * time.Microsecond)
	t.anims = append(t.anims, anim2)

	mat2 := material.NewStandard(&math32.Color{1, 1, 1})
	mat2.AddTexture(tex2)
	mat2.SetOpacity(1)
	mat2.SetTransparent(true)
	s2 := graphic.NewSprite(2, 2, mat2)
	s2.SetPosition(2, 2, 0)
	a.Scene().Add(s2)

	// Creates texture 3 and animator
	tex3, err := texture.NewTexture2DFromImage(a.DirData() + "/images/explosion4.png")
	if err != nil {
		a.Log().Fatal("Error loading texture: %s", err)
	}
	anim3 := texture.NewAnimator(tex3, 40, 1)
	anim3.SetDispTime(2 * 16666 * time.Microsecond)
	t.anims = append(t.anims, anim3)

	mat3 := material.NewStandard(&math32.Color{1, 1, 1})
	mat3.AddTexture(tex3)
	mat3.SetOpacity(0.8)
	mat3.SetTransparent(true)
	s3 := graphic.NewSprite(3, 3, mat3)
	s3.SetPosition(-2, -2, 0)
	a.Scene().Add(s3)

	// Creates texture 4 and animator
	tex4, err := texture.NewTexture2DFromImage(a.DirData() + "/images/walksequence.png")
	if err != nil {
		a.Log().Fatal("Error loading texture: %s", err)
	}
	anim4 := texture.NewAnimator(tex4, 6, 5)
	anim4.SetDispTime(2 * 16333 * time.Microsecond)
	t.anims = append(t.anims, anim4)

	mat4 := material.NewStandard(&math32.Color{1, 1, 1})
	mat4.AddTexture(tex4)
	mat4.SetOpacity(1)
	mat4.SetTransparent(true)
	s4 := graphic.NewSprite(2, 2, mat4)
	s4.SetPosition(2, -2, 0)
	a.Scene().Add(s4)
}

// Update is called every frame.
func (t *SpriteAnim) Update(a *app.App, deltaTime time.Duration) {

	for _, anim := range t.anims {
		anim.Update(time.Now())
	}
}

// Cleanup is called once at the end of the demo.
func (t *SpriteAnim) Cleanup(a *app.App) {}

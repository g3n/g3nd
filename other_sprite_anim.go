package main

import (
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
	"time"
)

type SpriteAnim struct {
	anims []*texture.Animator
}

func init() {
	TestMap["other.sprite_anim"] = &SpriteAnim{}
}

func (t *SpriteAnim) Initialize(ctx *Context) {

    // Initialize list of animators
	t.anims = make([]*texture.Animator, 0)

    // Adds axis helper
	axis := graphic.NewAxisHelper(2)
	ctx.Scene.Add(axis)

    // Creates texture 1 and animator
	tex1, err := texture.NewTexture2DFromImage(ctx.DirData+"/images/explosion7.png")
	if err != nil {
		log.Fatal("Error loading texture: %s", err)
	}
	anim1 := texture.NewAnimator(tex1, 8, 8)
	anim1.SetDispTime(16666 * time.Microsecond)
	anim1.SetMaxCycles(4)
	t.anims = append(t.anims, anim1)

	mat1 := material.NewStandard(&math32.Color{1, 1, 1})
	mat1.AddTexture(tex1)
	mat1.SetOpacity(1)
	s1 := graphic.NewSprite(2, 2, mat1)
	s1.SetPosition(-2, 2, 0)
	ctx.Scene.Add(s1)

    // Creates texture 2 and animator
	tex2, err := texture.NewTexture2DFromImage(ctx.DirData+"/images/explosion3.png")
	if err != nil {
		log.Fatal("Error loading texture: %s", err)
	}
	anim2 := texture.NewAnimator(tex2, 4, 4)
	anim2.SetDispTime(4 * 16666 * time.Microsecond)
	t.anims = append(t.anims, anim2)

	mat2 := material.NewStandard(&math32.Color{1, 1, 1})
	mat2.AddTexture(tex2)
	mat2.SetOpacity(1)
	s2 := graphic.NewSprite(2, 2, mat2)
	s2.SetPosition(2, 2, 0)
	ctx.Scene.Add(s2)

    // Creates texture 3 and animator
	tex3, err := texture.NewTexture2DFromImage(ctx.DirData+"/images/explosion4.png")
	if err != nil {
		log.Fatal("Error loading texture: %s", err)
	}
	anim3 := texture.NewAnimator(tex3, 40, 1)
	anim3.SetDispTime(2 * 16666 * time.Microsecond)
	t.anims = append(t.anims, anim3)

	mat3 := material.NewStandard(&math32.Color{1, 1, 1})
	mat3.AddTexture(tex3)
	mat3.SetOpacity(0.8)
	s3 := graphic.NewSprite(3, 3, mat3)
	s3.SetPosition(-2, -2, 0)
	ctx.Scene.Add(s3)

    // Creates texture 4 and animator
	tex4, err := texture.NewTexture2DFromImage(ctx.DirData+"/images/walksequence.png")
	if err != nil {
		log.Fatal("Error loading texture: %s", err)
	}
	anim4 := texture.NewAnimator(tex4, 6, 5)
	anim4.SetDispTime(2 * 16333 * time.Microsecond)
	t.anims = append(t.anims, anim4)

	mat4 := material.NewStandard(&math32.Color{1, 1, 1})
	mat4.AddTexture(tex4)
	mat4.SetOpacity(1)
	s4 := graphic.NewSprite(2, 2, mat4)
	s4.SetPosition(2, -2, 0)
	ctx.Scene.Add(s4)
}

func (t *SpriteAnim) Render(ctx *Context) {

	for _, anim := range t.anims {
		anim.Update(ctx.Time)
	}
}



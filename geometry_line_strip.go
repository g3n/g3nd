package main

import (
    "github.com/g3n/engine/gls"
    "github.com/g3n/engine/math32"
    "github.com/g3n/engine/graphic"
    "github.com/g3n/engine/material"
    "github.com/g3n/engine/geometry"
)


func init() {
    TestMap["geometry.line_strip"] = &LineStrip{}
}

type LineStrip struct {}

func (t *LineStrip) Initialize(ctx *Context) {

    // Creates geometry
    geom := geometry.NewGeometry()
    vertices := math32.NewArrayF32(0,32)
    vertices.Append(
         0.0,   0.0,  0.0,
         0.5,   0.0,  0.0,
         0.0,   0.0, -0.5,
         0.0,   0.5,  0.0,
         0.0,  -0.5,  0.0,
         0.0,   0.0, -0.5,
        -0.5,   0.0,  0.0,
         0.0,   0.0,  0.0,
     )
    colors := math32.NewArrayF32(0,32)
    colors.Append(
         1.0,   0.0,  0.0,
         1.0,   0.0,  0.0,
         0.0,   1.0,  0.0,
         0.0,   1.0,  0.0,
         0.0,   0.0,  1.0,
         0.0,   0.0,  1.0,
         0.0,   1.0,  1.0,
         0.0,   1.0,  1.0,
    )
    geom.AddVBO(gls.NewVBO().AddAttrib("VertexPosition", 3).SetBuffer(vertices))
    geom.AddVBO(gls.NewVBO().AddAttrib("VertexColor", 3).SetBuffer(colors))

    // Creates basic material
    mat := material.NewBasic()
    mat.SetLineWidth(2.0)

    // Creates line strip with the specified geometry and material
    lines := graphic.NewLineStrip(geom, mat)
    ctx.Scene.Add(lines)
}

func (t *LineStrip) Render(ctx *Context) {
}



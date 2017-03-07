package main

import (
    "github.com/g3n/engine/graphic"
)


func init() {
    TestMap["helper.axis"] = &AxisHelper{}
}

type AxisHelper struct {}


func (t *AxisHelper) Initialize(ctx *Context) {

    ah := graphic.NewAxisHelper(1.0)
    ctx.Scene.Add(ah)
}

func (t *AxisHelper) Render(ctx *Context) {
}


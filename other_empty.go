package main

import (
)


func init() {
    TestMap["other.empty"] = &Empty{}
}

type Empty struct {}


func (t *Empty) Initialize(ctx *Context) {
}

func (t *Empty) Render(ctx *Context) {
}



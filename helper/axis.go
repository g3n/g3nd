package helper

import (
	"github.com/g3n/engine/graphic"
	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/demos"
)

func init() {
	demos.Map["helper.axis"] = &AxisHelper{}
}

type AxisHelper struct{}

func (t *AxisHelper) Initialize(a *app.App) {

	ah := graphic.NewAxisHelper(1.0)
	a.Scene().Add(ah)
}

func (t *AxisHelper) Render(a *app.App) {
}

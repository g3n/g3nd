package helper

import (
	"github.com/g3n/engine/graphic"
	"github.com/g3n/g3nd/demos"
	"github.com/g3n/g3nd/g3nd"
)

func init() {
	demos.Map["helper.axis"] = &AxisHelper{}
}

type AxisHelper struct{}

func (t *AxisHelper) Initialize(app *g3nd.App) {

	ah := graphic.NewAxisHelper(1.0)
	app.Scene().Add(ah)
}

func (t *AxisHelper) Render(app *g3nd.App) {
}

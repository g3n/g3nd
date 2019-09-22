package helper

import (
	"github.com/g3n/engine/util/helper"
	"github.com/g3n/g3nd/app"
	"time"
)

func init() {
	app.DemoMap["helper.axis"] = &AxisHelper{}
}

type AxisHelper struct{}

// Start is called once at the start of the demo.
func (t *AxisHelper) Start(a *app.App) {

	ah := helper.NewAxes(1.0)
	a.Scene().Add(ah)
}

// Update is called every frame.
func (t *AxisHelper) Update(a *app.App, deltaTime time.Duration) {}

// Cleanup is called once at the end of the demo.
func (t *AxisHelper) Cleanup(a *app.App) {}

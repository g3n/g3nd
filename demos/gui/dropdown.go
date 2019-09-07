package gui

import (
	"fmt"
	"time"

	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/gui/assets/icon"
	"github.com/g3n/g3nd/app"
)

func init() {
	app.DemoMap["gui.dropdown"] = &GuiDropdown{}
}

type GuiDropdown struct{}

// Start is called once at the start of the demo.
func (t *GuiDropdown) Start(a *app.App) {

	// Show and enable demo panel
	a.DemoPanel().SetRenderable(true)
	a.DemoPanel().SetEnabled(true)

	// DropDown simple
	dd1 := gui.NewDropDown(150, gui.NewImageLabel("init"))
	dd1.SetPosition(10, 10)
	a.DemoPanel().Add(dd1)
	for i := 1; i <= 10; i++ {
		dd1.Add(gui.NewImageLabel(fmt.Sprintf("item %2d", i)))
	}
	dd1.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		a.Log().Debug("DropDown 1 OnChange")
	})
	dd1.SetSelected(dd1.ItemAt(0))

	// DropDown with icons
	icons := []string{
		icon.ArrowBack,
		icon.ArrowDownward,
		icon.ArrowDropDown,
		icon.ArrowDropDownCircle,
		icon.ArrowDropUp,
		icon.ArrowForward,
		icon.ArrowUpward,
	}
	model := gui.NewImageLabel("")
	dd2 := gui.NewDropDown(150, model)
	dd2.SetPosition(100, dd1.Position().Y+dd1.Height()+50)
	a.DemoPanel().Add(dd2)
	for i := 1; i <= 10; i++ {
		item := gui.NewImageLabel(fmt.Sprintf("item %2d", i))
		item.SetIcon(string(icons[i%len(icons)]))
		dd2.Add(item)
	}
	dd2.SetSelected(dd2.ItemAt(1))

	// DropDown with images
	images := []string{
		"ok.png",
	}
	model = gui.NewImageLabel("")
	dd3 := gui.NewDropDown(150, model)
	dd3.SetPosition(200, dd2.Position().Y+dd2.Height()+50)
	a.DemoPanel().Add(dd3)
	for i := 1; i <= 10; i++ {
		item := gui.NewImageLabel(fmt.Sprintf("item %2d", i))
		ifile := a.DirData() + "/images/" + images[i%len(images)]
		img, err := gui.NewImage(ifile)
		if err != nil {
			a.Log().Fatal("Error loading image:%s", err)
		}
		item.SetImage(img)
		dd3.Add(item)
	}
	dd3.SetSelected(dd3.ItemAt(2))
}

// Update is called every frame.
func (t *GuiDropdown) Update(a *app.App, deltaTime time.Duration) {}

// Cleanup is called once at the end of the demo.
func (t *GuiDropdown) Cleanup(a *app.App) {}

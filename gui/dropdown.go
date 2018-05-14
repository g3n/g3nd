package gui

import (
	"fmt"

	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/gui/assets/icon"
	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/demos"
)

func init() {
	demos.Map["gui.dropdown"] = &GuiDropdown{}
}

type GuiDropdown struct{}

func (t *GuiDropdown) Initialize(app *app.App) {

	// DropDown simple
	dd1 := gui.NewDropDown(150, gui.NewImageLabel("init"))
	dd1.SetPosition(10, 10)
	app.GuiPanel().Add(dd1)
	for i := 1; i <= 10; i++ {
		dd1.Add(gui.NewImageLabel(fmt.Sprintf("item %2d", i)))
	}
	dd1.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		app.Log().Debug("DropDown 1 OnChange")
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
	app.GuiPanel().Add(dd2)
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
	app.GuiPanel().Add(dd3)
	for i := 1; i <= 10; i++ {
		item := gui.NewImageLabel(fmt.Sprintf("item %2d", i))
		ifile := app.DirData() + "/images/" + images[i%len(images)]
		img, err := gui.NewImage(ifile)
		if err != nil {
			app.Log().Fatal("Error loading image:%s", err)
		}
		item.SetImage(img)
		dd3.Add(item)
	}
	dd3.SetSelected(dd3.ItemAt(2))
}

func (t *GuiDropdown) Render(app *app.App) {
}

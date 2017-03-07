package main

import (
    "fmt"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/graphic"
    "github.com/g3n/engine/gui/assets"
)


func init() {
	TestMap["gui.dropdown"] = &GuiDropdown{}
}

type GuiDropdown struct {}

func (t *GuiDropdown) Initialize(ctx *Context) {

	axis := graphic.NewAxisHelper(1)
    ctx.Scene.Add(axis)

    // DropDown simple
    dd1 := gui.NewDropDown(150, gui.NewImageLabel("init"))
	dd1.SetPosition(10, 10)
	ctx.Gui.Add(dd1)
    for i := 1; i <= 10; i++ {
        dd1.Add(gui.NewImageLabel(fmt.Sprintf("item %2d", i)))
    }
    dd1.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
        log.Debug("DropDown 1 OnChange") 
    })
    dd1.SetSelected(dd1.ItemAt(0))

    // DropDown with icons
    icons := []int{
	    assets.ArrowBack,
	    assets.ArrowDownward,
	    assets.ArrowDropDown,
	    assets.ArrowDropDownCircle,
	    assets.ArrowDropUp,
	    assets.ArrowForward,
	    assets.ArrowUpward,
    }
    model := gui.NewImageLabel("")
    dd2 := gui.NewDropDown(150, model)
	dd2.SetPosition(200, 10)
	ctx.Gui.Add(dd2)
    for i := 1; i <= 10; i++ {
        item := gui.NewImageLabel(fmt.Sprintf("item %2d", i))
        item.SetIcon(icons[i % len(icons)])
        dd2.Add(item)
    }
    dd2.SetSelected(dd2.ItemAt(1))

    // DropDown with images
    images := []string{
        "ok.png",
    }
    model = gui.NewImageLabel("")
    dd3 := gui.NewDropDown(150, model)
	dd3.SetPosition(400, 10)
	ctx.Gui.Add(dd3)
    for i := 1; i <= 10; i++ {
        item := gui.NewImageLabel(fmt.Sprintf("item %2d", i))
        ifile := ctx.DirData + "/images/" + images[i % len(images)]
        img, err := gui.NewImage(ifile)
        if err != nil {
            log.Fatal("Error loading image:%s", err)
        }
        item.SetImage(img)
        dd3.Add(item)
    }
    dd3.SetSelected(dd3.ItemAt(2))
}

func (t *GuiDropdown) Render(ctx *Context) {
}






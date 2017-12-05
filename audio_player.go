package main

import (
	"fmt"
	"github.com/g3n/engine/audio"
	"github.com/g3n/engine/audio/al"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/gui/assets/icon"
	"github.com/g3n/engine/math32"
	"time"
)

func init() {
	TestMap["audio.player"] = &AudioPlayer{}
}

type AudioPlayer struct {
	lastUpdate time.Time
	pc1        *PlayerControl
	pc2        *PlayerControl
	pc3        *PlayerControl
	pc4        *PlayerControl
}

func (t *AudioPlayer) Initialize(ctx *Context) {

	var err error
	t.pc1, err = NewPlayerControl(ctx, "bomb1.wav")
	t.pc1.player.SetLooping(false)
	if err != nil {
		log.Fatal("%s", err)
	}
	t.pc1.SetPosition(40, 10)
	ctx.Gui.Add(t.pc1)

	t.pc2, err = NewPlayerControl(ctx, "Vivaldi1.wav")
	if err != nil {
		log.Fatal("%s", err)
	}
	t.pc2.SetPosition(40, t.pc1.Position().Y+t.pc1.Height()+20)
	ctx.Gui.Add(t.pc2)

	t.pc3, err = NewPlayerControl(ctx, "bomb2.ogg")
	t.pc3.player.SetLooping(false)
	if err != nil {
		log.Fatal("%s", err)
	}
	t.pc3.SetPosition(40, t.pc2.Position().Y+t.pc2.Height()+40)
	ctx.Gui.Add(t.pc3)

	t.pc4, err = NewPlayerControl(ctx, "Bach1.ogg")
	if err != nil {
		log.Fatal("%s", err)
	}
	t.pc4.SetPosition(40, t.pc3.Position().Y+t.pc3.Height()+20)
	ctx.Gui.Add(t.pc4)
}

func (t *AudioPlayer) Render(ctx *Context) {

	if time.Now().Sub(t.lastUpdate) < 100*time.Millisecond {
		return
	}
	t.pc1.UpdateTime()
	t.pc2.UpdateTime()
	t.pc3.UpdateTime()
	t.pc4.UpdateTime()
	t.lastUpdate = time.Now()
}

type PlayerControl struct {
	gui.Panel   // Embedded panel
	player      *audio.Player
	title       *gui.Label
	sliderTime  *gui.Slider
	sliderGain  *gui.Slider
	sliderPitch *gui.Slider
}

func NewPlayerControl(ctx *Context, filename string) (*PlayerControl, error) {

	// Creates player
	player, err := audio.NewPlayer(ctx.DirData + "/audio/" + filename)
	if err != nil {
		return nil, err
	}

	pc := new(PlayerControl)
	pc.player = player

	pc.Panel.Initialize(300, 90)
	pc.SetBorders(1, 1, 1, 1)
	pc.SetColor(&math32.Color{0.5, 0.5, 0.5})
	pc.SetPaddings(4, 4, 4, 4)

	// External layout
	layout := gui.NewHBoxLayout()
	layout.SetSpacing(4)
	pc.SetLayout(layout)

	// Panel for left group
	panLeft := gui.NewPanel(0, 0)
	panLeft.SetBorders(0, 0, 0, 0)
	panLeft.SetLayoutParams(&gui.HBoxLayoutParams{AlignV: gui.AlignHeight, Expand: 4.5})
	panLeftLayout := gui.NewVBoxLayout()
	panLeft.SetLayout(panLeftLayout)
	pc.Add(panLeft)

	// Panel for right group
	panRight := gui.NewPanel(0, 0)
	panRight.SetBorders(0, 0, 0, 0)
	panRight.SetPaddings(0, 0, 0, 4)
	panRight.SetLayoutParams(&gui.HBoxLayoutParams{AlignV: gui.AlignHeight, Expand: 1})
	panRightLayout := gui.NewHBoxLayout()
	panRightLayout.SetSpacing(4)
	panRight.SetLayout(panRightLayout)
	pc.Add(panRight)

	// File name
	pc.title = gui.NewLabel("title")
	pc.title.SetLayoutParams(&gui.VBoxLayoutParams{AlignH: gui.AlignCenter, Expand: 0})
	panLeft.Add(pc.title)

	// Slider for current time
	pc.sliderTime = gui.NewHSlider(0, 24)
	pc.sliderTime.SetLayoutParams(&gui.VBoxLayoutParams{AlignH: gui.AlignWidth, Expand: 0})
	pc.sliderTime.SetEnabled(false)
	sliderTimeStyles := gui.StyleDefault.Slider
	sliderTimeStyles.Disabled.FgColor.SetName("dodgerblue")
	pc.sliderTime.SetStyles(&sliderTimeStyles)
	panLeft.Add(pc.sliderTime)

	// Panel for control buttons
	bPanel := gui.NewPanel(0, 0)
	bPanel.SetBorders(0, 0, 0, 0)
	bPanelLayout := gui.NewHBoxLayout()
	bPanelLayout.SetSpacing(10)
	bPanelLayout.SetAlignH(gui.AlignWidth)
	bPanel.SetLayout(bPanelLayout)

	// Play button
	bPlay := gui.NewButton("Play")
	bPlay.SetIcon(icon.PlayArrow)
	bPlay.SetLayoutParams(&gui.HBoxLayoutParams{AlignV: gui.AlignCenter})
	bPlay.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		pc.player.Play()
	})
	bPanel.Add(bPlay)

	// Pause button
	bPause := gui.NewButton("Pause")
	bPause.SetIcon(icon.Pause)
	bPause.SetLayoutParams(&gui.HBoxLayoutParams{AlignV: gui.AlignCenter})
	bPause.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		pc.player.Pause()
	})
	bPanel.Add(bPause)

	// Stop button
	bStop := gui.NewButton("Stop")
	bStop.SetIcon(icon.Stop)
	bStop.SetLayoutParams(&gui.HBoxLayoutParams{AlignV: gui.AlignCenter})
	bStop.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		pc.player.Stop()
		pc.UpdateTime()
	})
	bPanel.Add(bStop)

	// Sets button panel layout and adds it to player panel
	bPanel.SetLayoutParams(&gui.VBoxLayoutParams{AlignH: gui.AlignWidth, Expand: 1})
	panLeft.Add(bPanel)

	// Gain Slider
	pc.sliderGain = gui.NewVSlider(20, 0)
	pc.sliderGain.SetLayoutParams(&gui.HBoxLayoutParams{AlignV: gui.AlignHeight, Expand: 0})
	pc.sliderGain.SetEnabled(true)
	pc.sliderGain.SetText("G")
	pc.sliderGain.SetValue(pc.player.Gain())
	pc.sliderGain.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		pc.player.SetGain(pc.sliderGain.Value())
	})
	sliderGainStyles := gui.StyleDefault.Slider
	sliderGainStyles.Normal.FgColor.SetName("darkseagreen")
	sliderGainStyles.Over.FgColor.SetName("darkseagreen").MultiplyScalar(1.1)
	pc.sliderGain.SetStyles(&sliderGainStyles)
	panRight.Add(pc.sliderGain)

	// Pitch Slider
	pc.sliderPitch = gui.NewVSlider(20, 0)
	pc.sliderPitch.SetLayoutParams(&gui.HBoxLayoutParams{AlignV: gui.AlignHeight, Expand: 0})
	pc.sliderPitch.SetEnabled(true)
	pc.sliderPitch.SetText("P")
	pc.sliderPitch.SetValue(pc.player.Pitch() / 2)
	pc.sliderPitch.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		pc.player.SetPitch(pc.sliderPitch.Value() * 2)
	})
	sliderPitchStyles := gui.StyleDefault.Slider
	sliderPitchStyles.Normal.FgColor.SetName("darkseagreen")
	sliderPitchStyles.Over.FgColor.SetName("darkseagreen").MultiplyScalar(1.1)
	pc.sliderPitch.SetStyles(&sliderPitchStyles)
	panRight.Add(pc.sliderPitch)

	pc.title.SetText(filename)
	return pc, nil
}

func (pc *PlayerControl) Dispose() {

	pc.Panel.Dispose()
	pc.player.Dispose()
}

func (pc *PlayerControl) UpdateTime() {

	if pc.player.State() != al.Playing {
		return
	}
	msg := fmt.Sprintf("%1.1f / %1.1f", pc.player.CurrentTime(), pc.player.TotalTime())
	pc.sliderTime.SetText(msg)
	pc.sliderTime.SetValue(float32(pc.player.CurrentTime() / pc.player.TotalTime()))
}

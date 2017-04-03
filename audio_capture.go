package main

import (
	"github.com/g3n/engine/audio/al"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
)

func init() {
	TestMap["audio.capture"] = &AudioCapture{}
}

type AudioCapture struct {
	buf     []byte    // buffer with mono16 samples
	samples []float32 // buffer with normalized float32 samples
	chart   *gui.Chart
	gr      *gui.Graph
}

const (
	capRate    = 48000 // Capture sample rate
	capSamples = 960   // Number of samples to show (~20ms)
)

func (t *AudioCapture) Initialize(ctx *Context) {

	// Try to open default capture device
	dev, err := al.CaptureOpenDevice("", capRate, al.FormatMono16, 2*capSamples)
	if err != nil {
		l := gui.NewLabel("Error opening default capture device")
		l.SetFontSize(22)
		px := (ctx.Gui.Width() - l.Width()) / 2
		py := ctx.Gui.Height() / 2
		l.SetPosition(px, py)
		log.Error("%s", err)
		ctx.Gui.Add(l)
		return
	}

	// Save capture device so main program can close it.
	ctx.CapDev = dev

	// Creates Chart panel
	t.chart = gui.NewChart(500, 300)
	t.chart.SetMargins(10, 10, 10, 10)
	t.chart.SetBorders(2, 2, 2, 2)
	t.chart.SetBordersColor(&math32.Black)
	t.chart.SetColor(&math32.White)
	t.chart.SetPaddings(2, 2, 2, 2)
	t.chart.SetPosition(0, 0)
	t.chart.SetTitle("Captured audio waveform", 15)
	// X scale
	t.chart.SetFormatX("%2.1f")
	t.chart.SetRangeX(0, 0, capSamples)
	// Y scale
	t.chart.SetFormatY("%2.1f")
	t.chart.SetScaleY(5, &math32.Color{0.8, 0.8, 0.8})
	t.chart.SetFontSizeY(13)
	t.chart.SetRangeY(-1.0, 1.0)
	ctx.Gui.Add(t.chart)

	// Adds line graph
	t.gr = t.chart.AddLineGraph(&math32.Color{0, 0, 1}, nil)

	// Creates buffers
	t.buf = make([]byte, 2*capSamples)
	t.samples = make([]float32, capSamples)

	// Start capturing samples
	al.CaptureStart(ctx.CapDev)
}

func (t *AudioCapture) Render(ctx *Context) {

	// If device was not created successfully, nothing to do
	if ctx.CapDev == nil {
		return
	}

	// Get how many samples were captured
	values := make([]int32, 1)
	al.CtxGetIntegerv(ctx.CapDev, al.CtxCaptureSamples, values)
	if values[0] < capSamples {
		return
	}

	// Read captured samples, builds normalized float32 buffer and plot data
	al.CaptureSamples(ctx.CapDev, t.buf, capSamples)
	for i := 0; i < capSamples; i++ {
		s := int16(t.buf[i*2]) + int16(t.buf[i*2+1])*256
		t.samples[i] = float32(s) / 32767
	}
	t.gr.SetData(t.samples)
}

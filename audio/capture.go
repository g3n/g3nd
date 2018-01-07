package audio

import (
	"github.com/g3n/engine/audio/al"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/demos"
)

func init() {
	demos.Map["audio.capture"] = &AudioCapture{}
}

type AudioCapture struct {
	buf     []byte    // buffer with mono16 samples
	samples []float32 // buffer with normalized float32 samples
	chart   *gui.Chart
	gr      *gui.Graph
	capDev  *al.Device // Audio capture device
}

const (
	capRate    = 48000 // Capture sample rate
	capSamples = 960   // Number of samples to show (~20ms)
)

func (t *AudioCapture) Initialize(a *app.App) {

	// Try to open default capture device
	dev, err := al.CaptureOpenDevice("", capRate, al.FormatMono16, 2*capSamples)
	if err != nil {
		l := gui.NewLabel("Error opening default capture device")
		l.SetFontSize(22)
		px := (a.GuiPanel().Width() - l.Width()) / 2
		py := a.GuiPanel().Height() / 2
		l.SetPosition(px, py)
		a.Log().Error("%s", err)
		a.GuiPanel().Add(l)
		return
	}

	// Adds function to close capture device
	t.capDev = dev
	a.AddFinalizer(func() {
		al.CaptureStop(t.capDev)
		al.CaptureCloseDevice(t.capDev)
		a.Log().Debug("Audio capture device closed")
	})

	// Creates chart panel
	t.chart = gui.NewChart(500, 300)
	t.chart.SetMargins(10, 10, 10, 10)
	t.chart.SetBorders(2, 2, 2, 2)
	t.chart.SetBordersColor(math32.NewColor("black"))
	t.chart.SetColor(math32.NewColor("white"))
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
	a.GuiPanel().Add(t.chart)

	// Adds line graph
	t.gr = t.chart.AddLineGraph(&math32.Color{0, 0, 1}, nil)

	// Creates buffers
	t.buf = make([]byte, 2*capSamples)
	t.samples = make([]float32, capSamples)

	// Start capturing audio samples
	al.CaptureStart(t.capDev)
}

func (t *AudioCapture) Render(a *app.App) {

	// If device was not created successfully, nothing to do
	if t.capDev == nil {
		return
	}

	// Get how many samples were captured
	values := make([]int32, 1)
	al.CtxGetIntegerv(t.capDev, al.CtxCaptureSamples, values)
	if values[0] < capSamples {
		return
	}

	// Read captured samples, builds normalized float32 buffer and plot data
	al.CaptureSamples(t.capDev, t.buf, capSamples)
	for i := 0; i < capSamples; i++ {
		s := int16(t.buf[i*2]) + int16(t.buf[i*2+1])*256
		t.samples[i] = float32(s) / 32767
	}
	t.gr.SetData(t.samples)
}

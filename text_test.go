package cache2d

import (
	"image/color"
	"os"
	"runtime"
	"testing"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/redstarcoder/draw2d"
	"github.com/redstarcoder/draw2d/draw2dgl"
)

var (
	width, height int
	mx, my        int
	font          draw2d.FontData
	gc            draw2d.GraphicContext
)

const FUNTEXT = "qwertyuiopasdfghjklzxcvbnm1234567890~!@#$%^&*()_+{}|:'<>?/✪"

func BenchmarkFillStringAt(b *testing.B) {
	b.StopTimer()
	var err error
	width, height = 800, 800
	window, err := glfw.CreateWindow(width, height, "BenchmarkFillStringAt", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	glfw.SwapInterval(0)

	reshape(window, width, height)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		displayString()
		b.StopTimer()
		glfw.PollEvents()
	}
	window.Destroy()
}

func BenchmarkFillStringAtCached(b *testing.B) {
	b.StopTimer()
	width, height = 800, 800
	window, err := glfw.CreateWindow(width, height, "BenchmarkFillStringAtCached", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	glfw.SwapInterval(0)

	reshape(window, width, height)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		displayStringCached()
		b.StopTimer()
		glfw.PollEvents()
	}
	window.Destroy()
}

func TestMain(m *testing.M) {
	r := m.Run()
	glfw.Terminate()
	os.Exit(r)
}

func reshape(window *glfw.Window, w, h int) {
	gl.ClearColor(1, 1, 1, 1)
	/* Establish viewing area to cover entire window. */
	gl.Viewport(0, 0, int32(w), int32(h))
	/* PROJECTION Matrix mode. */
	gl.MatrixMode(gl.PROJECTION)
	/* Reset project matrix. */
	gl.LoadIdentity()
	/* Map abstract coords directly to window coords. */
	gl.Ortho(0, float64(w), 0, float64(h), -1, 1)
	/* Invert Y axis so increasing Y goes down. */
	gl.Scalef(1, -1, 1)
	/* Shift origin up to upper-left corner. */
	gl.Translatef(0, float32(-h), 0)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Disable(gl.DEPTH_TEST)
	width, height = w, h
	/* Recreate graphic context with new width & height. */
	gc = draw2dgl.NewGraphicContext(width, height)
	gc.SetFontData(draw2d.FontData{
		Name:   "luxi",
		Family: draw2d.FontFamilyMono,
		Style:  draw2d.FontStyleBold | draw2d.FontStyleItalic})
	gc.SetFontSize(14)
}

func displayString() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	// Display FUNTEXT
	gl.LineWidth(1)
	gc.SetFillColor(color.RGBA{0, 0, 0, 0xff})
	gc.FillStringAt(FUNTEXT, 10, gc.GetFontSize()+10)

	gl.Flush() /* Single buffered, so needs a flush. */
}

func displayStringCached() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	// Display FUNTEXT
	gl.LineWidth(1)
	gc.SetFillColor(color.RGBA{0, 0, 0, 0xff})
	FillStringByGlyph(gc, FUNTEXT, 10, gc.GetFontSize()+10)

	gl.Flush() /* Single buffered, so needs a flush. */
}

func init() {
	runtime.LockOSThread()
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	err = gl.Init()
	if err != nil {
		panic(err)
	}
	draw2d.SetFontFolder("resources/font")
}

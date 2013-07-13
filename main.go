package main

import (
	gl "github.com/chsc/gogl/gl21"
	"github.com/runningwild/glop/gin"
	"github.com/runningwild/glop/gos"
	"github.com/runningwild/glop/gui"
	"github.com/runningwild/glop/render"
	"github.com/runningwild/glop/system"
	"os"
	"path/filepath"
	"runtime"
)

func init() {
	// Required for Darwin.
	runtime.LockOSThread()
}

func initWindow(sys system.System, width int, height int) {
	sys.CreateWindow(10, 10, width, height)
	sys.EnableVSync(true)
	err := gl.Init()
	if err != nil {
		panic(err)
	}
	gl.Ortho(gl.Double(0), gl.Double(width), gl.Double(0), gl.Double(height),
		1000, -1000)
}

func loadDictionary(fontName string) *gui.Dictionary {
	f, err := os.Open(filepath.Join("..", "data", "skia.gob"))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	dict, err := gui.LoadDictionary(f)
	if err != nil {
		panic(err)
	}
	return dict
}

func main() {
	runtime.GOMAXPROCS(2)
	sys := system.Make(gos.GetSystemInterface())
	sys.Startup()

	render.Init()
	render.Queue(func() {
		initWindow(sys, 800, 600)
	})

	font := loadDictionary("skia.gob")

	for true {
		sys.Think()
		render.Queue(func() {
			gl.Color4ub(255, 255, 255, 255)
			gl.Begin(gl.QUADS)
			gl.Vertex2d(100, 100)
			gl.Vertex2d(500, 100)
			gl.Vertex2d(500, 150)
			gl.Vertex2d(100, 150)
			gl.End()
			font.RenderString("TEST", 100, 100, 0, 100, gui.Left)
			sys.SwapBuffers()
		})
		render.Purge()
		if gin.In().GetKey(gin.AnyEscape).FramePressCount() > 0 {
			break
		}
	}
}

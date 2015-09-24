package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	gl "github.com/go-gl/gl/v3.3-core/gl"
	"github.com/runningwild/glop/gin"
	"github.com/runningwild/glop/gos"
	"github.com/runningwild/glop/render"
	"github.com/runningwild/glop/system"
	"github.com/runningwild/glop/text"
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
	// gl.Ortho(gl.Double(0), gl.Double(width), gl.Double(0), gl.Double(height),
	// 	1000, -1000)
}

var log *os.File

func init() {
	var err error
	log, err = os.Create("/Users/jwills/code/src/github.com/dgthunder/glomple/log.txt")
	if err != nil {
		panic(err)
	}
}

func loadDictionary(fontName string) *text.Dictionary {
	fmt.Fprintf(log, "path: %v\n", filepath.Join(os.Args[0], "..", "..", "skia.dict"))
	f, err := os.Open(filepath.Join(os.Args[0], "..", "..", "skia.dict"))
	fmt.Fprintf(log, "err: %v\n", err)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Fprintf(log, "err: %v\n", err)
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(log, "recover: %v\n", r)
		}
	}()
	dict, err := text.LoadDictionary(f)
	if err != nil {
		panic(err)
	}
	return dict
}

func main() {
	sys := system.Make(gos.GetSystemInterface())
	sys.Startup()

	render.Init()
	render.Queue(func() {
		initWindow(sys, 800, 600)
	})

	font := loadDictionary("skia.dict")
	fmt.Fprintf(log, "Font: %v", font)
	for true {
		sys.Think()
		render.Queue(func() {
			// gl.Color4ub(0, 255, 255, 255)
			// gl.Begin(gl.QUADS)
			// gl.Vertex2d(100, 100)
			// gl.Vertex2d(500, 100)
			// gl.Vertex2d(500, 150)
			// gl.Vertex2d(100, 150)
			// gl.End()
			font.SetFontColor(1, 1, 1)
			font.RenderString("TEST", 100, 100, 100)
			sys.SwapBuffers()
		})
		render.Purge()
		if gin.In().GetKey(gin.AnyEscape).FramePressCount() > 0 {
			break
		}
	}
}

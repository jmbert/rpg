package graphics

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

var window *sdl.Window
var renderer *sdl.Renderer

var Width = 0
var Height = 0

func Initialise(width, height int32, wflags, rflags uint32, title string) {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		log.Fatalln(err)
	}

	win, err := sdl.CreateWindow(title, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, width, height, wflags)
	if err != nil {
		log.Fatalln(err)
	}

	ren, err := sdl.CreateRenderer(win, -1, rflags)
	if err != nil {
		log.Fatalln(err)
	}

	ren.SetDrawColor(0, 0, 0, 0)
	ren.Clear()

	window = win
	renderer = ren

	Width = int(width)
	Height = int(height)
}

func Destroy() {
	window.Destroy()
	renderer.Destroy()

	sdl.Quit()
}

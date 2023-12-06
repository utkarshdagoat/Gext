package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var (
	winTitle            string = "Go-SDL2 Render"
	winWidth, winHeight int32  = 800, 600
)

func run() int {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var surface *sdl.Surface
	var texture *sdl.Texture

	white := sdl.Color{255, 255, 255, 255}

	font_ttf, err := ttf.OpenFont("./JetBrainsMono-Regular.ttf", 24)
	if err != nil {
		fmt.Fprintf(os.Stderr, "FONT:SDL Error: %s \n", err)
		os.Exit(2)
	}

	surface, err = font_ttf.RenderUTF8Solid("Hello,World", white)
	if err != nil {
		fmt.Fprintf(os.Stderr, "SURFACE: SDL Error: %s \n", err)
		os.Exit(2)
	}

	rect := sdl.Rect{20, 20, *(&surface.W), *(&surface.H)}

	defer surface.Free()

	window, err = sdl.CreateWindow(winTitle, 0, 0,
		winWidth, winHeight, sdl.WINDOW_RESIZABLE)
	if err != nil {
		fmt.Fprintf(os.Stderr, "WINDOW: SDL Error: %s \n", err)
		os.Exit(2)
	}

	defer window.Destroy()

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	if err != nil {
		fmt.Fprintf(os.Stderr, "RENDERER: SDL Error: %s \n", err)
		os.Exit(2)
	}

	defer renderer.Destroy()

	running := true

	texture, err = renderer.CreateTextureFromSurface(surface)

	if err != nil {
		fmt.Fprintf(os.Stderr, "TEXTURE: SDL Error: %s \n", err)
		os.Exit(2)
	}
	defer texture.Destroy()

	fmt.Printf("TextEvent: %d \n", sdl.TEXTINPUT)

	sdl.StartTextInput()
	var textInput string

	for running {
		// I will leave this here
		// Event is a fucking generic empty interface spent ages trying to figure it out
		// ceretified retard moment didn't know how interfaces work

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.TextInputEvent:
				e := event.(*sdl.TextInputEvent)
				textInput += e.GetText()
				fmt.Println(textInput)
				if texture != nil {
					println("Here")
					texture.Destroy()
				}
				surface, err = font_ttf.RenderUTF8Solid(textInput, white)
				if err != nil {
					fmt.Fprintf(os.Stderr, "SURFACE WHILE EDITING: SDL Error: %s \n", err)
					os.Exit(2)
				}
				texture, err = renderer.CreateTextureFromSurface(surface)
				if err != nil {
					fmt.Fprintf(os.Stderr, "TEXTURE WHILE EDITING: SDL Error: %s \n", err)
					os.Exit(2)
				}

				rect.W = surface.W
				rect.H = surface.H

			}
		}

		renderer.Clear()
		renderer.Copy(texture, nil, &rect)
		renderer.Present()
		sdl.Delay(16)

	}
	sdl.StopTextInput()

	return 0
}

func main() {
	// SDL initilaizing
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Fprint(os.Stderr, "INIT:SDL ERROR: %s \n", err)
		os.Exit(2)
	}

	err = ttf.Init()
	if err != nil {
		fmt.Fprint(os.Stderr, "TTF INIT:SDL ERROR: %s \n", err)
		os.Exit(2)
	}

	code := run()
	sdl.Quit()
	ttf.Quit()
	os.Exit(code)
}

package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"github.com/andyaspel/pixl/apptypes"
	"github.com/andyaspel/pixl/swatch"
	"github.com/andyaspel/pixl/ui"
)

func main() {
	init := app.New()
	window := init.NewWindow("Pixl")
	state := apptypes.State{
		BrushColor:    color.NRGBA{255, 255, 255, 255},
		SwatchSeleted: 0,
	}
	appInit := ui.AppInit{
		Window:   window,
		State:    &state,
		Swatches: make([]*swatch.Swatch, 0, 32),
	}
	ui.Setup(&appInit)
	appInit.Window.ShowAndRun()
}

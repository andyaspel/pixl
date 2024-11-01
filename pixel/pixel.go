package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"github.com/andyaspel/pixel/apptypes"
	"github.com/andyaspel/pixel/swatches"
	"github.com/andyaspel/pixel/ui"
)

func main() {
	init := app.New()
	window := init.NewWindow("pixel")
	state := apptypes.State{
		BrushColor:    color.NRGBA{255, 255, 255, 255},
		SwatchSeleted: 0,
	}
	appInit := ui.AppInit{
		Window:   window,
		State:    &state,
		Swatches: make([]*swatches.Swatch, 0, 64),
	}
	ui.Setup(&appInit)
	appInit.Window.ShowAndRun()
}

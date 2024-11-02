package ui

import (
	"fyne.io/fyne/v2"
	"github.com/andyaspel/pixl/apptypes"
	"github.com/andyaspel/pixl/swatch"
)

type AppInit struct {
	Window   fyne.Window
	State    *apptypes.State
	Swatches []*swatch.Swatch
}

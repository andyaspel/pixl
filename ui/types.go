package ui

import (
	"fyne.io/fyne/v2"
	"github.com/andyaspel/pixl/apptypes"
	"github.com/andyaspel/pixl/swatches"
)

type AppInit struct {
	Window   fyne.Window
	State    *apptypes.State
	Swatches []*swatches.Swatch
}

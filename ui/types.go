package ui

import (
	"fyne.io/fyne/v2"
	"github.com/andyaspel/pixel/apptypes"
	"github.com/andyaspel/pixel/swatches"
)

type AppInit struct {
	Window   fyne.Window
	State    *apptypes.State
	Swatches []*swatches.Swatch
}

package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/andyaspel/pixel/swatches"
)

func BuildSwatches(app *AppInit) *fyne.Container {
	swatchContainer := make([]fyne.CanvasObject, 0, 64)
	for i := 0; i < cap(app.Swatches); i++ {
		initColor := color.NRGBA{255, 255, 255, 255}
		s := swatches.NewSwatch(app.State, initColor, i, func(s *swatches.Swatch) {
			for j := 0; j < len(app.Swatches); j++ {
				app.Swatches[j].Selected = false
				swatchContainer[j].Refresh()
			}
			app.State.SwatchSeleted = s.SwatchIndex
			app.State.BrushColor = s.Color
		})
		if i == 0 {
			s.Selected = true
			app.State.SwatchSeleted = 0
			s.Refresh()
		}
		app.Swatches = append(app.Swatches, s)
		swatchContainer = append(swatchContainer, s)
	}
	return container.NewGridWrap(fyne.NewSize(20, 20), swatchContainer...)
}

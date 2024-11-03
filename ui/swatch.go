package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/andyaspel/pixl/swatch"
)

func BuildSwatches(app *AppInit) *fyne.Container {
	swatchContainer := make([]fyne.CanvasObject, 0, 32)
	for i := 0; i < cap(app.Swatches); i++ {
		initColor := color.NRGBA{100, 0, 0, 255}
		s := swatch.NewSwatch(app.State, initColor, i, func(s *swatch.Swatch) {
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
	return container.NewGridWrap(fyne.NewSize(50, 50), swatchContainer...)
}

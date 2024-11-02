package swatches

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/andyaspel/pixl/apptypes"
)

type Swatch struct {
	widget.BaseWidget
	Selected     bool
	Color        color.Color
	SwatchIndex  int
	clickHandler func(s *Swatch)
}

func (s *Swatch) SetColor(c color.Color) {
	s.Color = c
	s.Refresh()
}

func NewSwatch(state *apptypes.State, color color.Color, swatchIndex int, clickHandler func(s *Swatch)) *Swatch {
	swatch := &Swatch{
		Selected:     false,
		Color:        color,
		SwatchIndex:  swatchIndex,
		clickHandler: clickHandler,
	}

	swatch.ExtendBaseWidget(swatch)
	return swatch
}

func (s *Swatch) CreateSwatch() fyne.WidgetRenderer {
	square := canvas.NewRectangle(s.Color)
	objects := []fyne.CanvasObject{square}
	return &SwatchRednerer{
		pixel:   *square,
		objects: objects,
		parent:  s,
	}
}
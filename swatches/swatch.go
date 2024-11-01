package swatches

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/andyaspel/pixel/apptypes"
	"github.com/andyaspel/pixel/swatches"
)

type Swatch struct {
	widget.BaseWidget
	Selected     bool
	Color        color.Color
	SwatchIndex  int
	clickHandler func(swatches *Swatch)
}

func (s *Swatch) SetColor(c color.Color) {
	s.Color = c
	s.Refresh()
}

func NewSwatch(state *apptypes.State, color color.Color, index int, onClick func(s *Swatch)) *Swatch {
	s := &Swatch{
		Selected:     false,
		Color:        color,
		SwatchIndex:  index,
		clickHandler: onClick,
	}
	swatches.ExtendBaseWidget(s)
	return s
}

func (s *Swatch) CreateSwatch() fyne.WidgetRenderer {
	square := canvas.NewRectangle(s.Color)
	objects := []fyne.CanvasObject{square}
	return &SwatchRednerer{
		square:  *square,
		objects: objects,
		parent:  s,
	}
}

package swatches

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type SwatchRednerer struct {
	square  canvas.Rectangle
	objects []fyne.CanvasObject
	parent  *Swatch
}

func (r *SwatchRednerer) MinSize() fyne.Size {
	return r.square.MinSize()
}

func (r *SwatchRednerer) Layout(size fyne.Size) {
	r.objects[0].Resize(size)
}

func (r *SwatchRednerer) Refresh() {
	r.Layout(fyne.NewSize(20, 20))
	r.square.FillColor = r.parent.Color
	if r.parent.Selected {
		r.square.StrokeWidth = 3
		r.square.StrokeColor = color.NRGBA{255, 255, 255, 255}
		r.objects[0] = &r.square
	} else {
		r.square.StrokeWidth = 0
		r.objects[0] = &r.square
	}
	canvas.Refresh(r.parent)
}

func (r *SwatchRednerer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *SwatchRednerer) Destroy() {}
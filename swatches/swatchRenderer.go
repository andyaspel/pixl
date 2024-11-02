package swatches

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type SwatchRednerer struct {
	pixel   canvas.Rectangle
	objects []fyne.CanvasObject
	parent  *Swatch
}

func (r *SwatchRednerer) MinSize() fyne.Size {
	return r.pixel.MinSize()
}

func (r *SwatchRednerer) Layout(size fyne.Size) {
	r.objects[0].Resize(size)
}

func (r *SwatchRednerer) Refresh() {
	r.Layout(fyne.NewSquareSize(20))
	r.pixel.FillColor = r.parent.Color
	if r.parent.Selected {
		r.pixel.StrokeWidth = 3
		r.pixel.StrokeColor = color.NRGBA{255, 255, 255, 255}
		r.objects[0] = &r.pixel
	} else {
		r.pixel.StrokeWidth = 0
		r.objects[0] = &r.pixel
	}
	canvas.Refresh(r.parent)
}

func (r *SwatchRednerer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *SwatchRednerer) Destroy() {}

package swatch

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type SwatchRenderer struct {
	pixel   canvas.Rectangle
	objects []fyne.CanvasObject
	parent  *Swatch
}

func (r *SwatchRenderer) MinSize() fyne.Size {
	return r.pixel.MinSize()
}

func (r *SwatchRenderer) Layout(size fyne.Size) {
	r.objects[0].Resize(size)
}

func (r *SwatchRenderer) Refresh() {
	r.Layout(fyne.NewSize(20, 20))
	r.pixel.FillColor = r.parent.Color
	if r.parent.Selected {
		r.pixel.StrokeWidth = 3
		r.pixel.StrokeColor = color.NRGBA{255, 255, 255, 255}
		r.objects[0] = &r.pixel
		fmt.Printf("Refresh-func if -\t%v\n", r)

	} else {
		r.pixel.StrokeWidth = 0
		r.objects[0] = &r.pixel
		fmt.Printf("Refresh-func else -\t%v\n", r)

	}
	canvas.Refresh(r.parent)
}

func (r *SwatchRenderer) Objects() []fyne.CanvasObject {
	fmt.Printf("Objects-func -\t%v\n", r)
	return r.objects
}

func (r *SwatchRenderer) Destroy() {}

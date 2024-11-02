package swatches

import (
	"fyne.io/fyne/v2/driver/desktop"
)

func (s *Swatch) MouseDown(e *desktop.MouseEvent) {
	s.clickHandler(s)
	s.Selected = true
	s.Refresh()
}

func (s *Swatch) MouseUp(e *desktop.MouseEvent) {}

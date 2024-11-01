package apptypes

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type BrushType = int
type PxCanvasConfig struct {
	CanvasArea     fyne.Size
	CanvasOffset   fyne.Position
	PxRows, PxCols int
	PxSize         int
}
type State struct {
	BrushColor    color.Color
	BrushType     int
	SwatchSeleted int
	FilePath      string
}

func (s *State) SetFilePath(path string) {
	s.FilePath = path
}

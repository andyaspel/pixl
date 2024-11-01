package ui

func Setup(a *AppInit) {
	swatchesBox := BuildSwatches(a)
	a.Window.SetContent(swatchesBox)
}

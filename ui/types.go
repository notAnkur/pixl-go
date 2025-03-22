package ui

import (
	"ankuranant.dev/pixl/apptype"
	"ankuranant.dev/pixl/swatch"
	"fyne.io/fyne/v2"
)

type AppInit struct {
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}

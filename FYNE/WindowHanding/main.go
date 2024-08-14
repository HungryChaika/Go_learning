package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	application := app.New()
	window := application.NewWindow("Hello world")

	window.SetContent(widget.NewLabel("Hello world!"))
	window.Show()

	window_another := application.NewWindow("Larger")
	window_another.SetContent(widget.NewLabel("More content"))
	window_another.Resize(fyne.NewSize(100, 100))
	window_another.Show()

	window_another.SetMaster()

	application.Run()
}

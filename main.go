package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Test")

	w.Resize(fyne.NewSize(480, 360))
	w.ShowAndRun()
	// hello := widget.NewLabel("Hello Fyne!")
	// w.SetContent(container.NewVBox(
	// 	hello,
	// 	widget.NewButton("Hi!", func() {
	// 		hello.SetText("Welcome :)")
	// 	}),
	// ))

	w.ShowAndRun()
}

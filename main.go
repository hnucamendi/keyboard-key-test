package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

const FULLKEYBOARDCOLUM int = 22

func buildKeyboard(ks int) *fyne.Container {
	t := canvas.NewText("Hello", color.White)
	ts := canvas.NewText("Hello", color.White)
	Keyboard := container.New(layout.NewGridLayout(ks), t, ts, t, ts)
	return Keyboard
}

func main() {
	a := app.New()
	w := a.NewWindow("Keyboard Tester")

	w.Resize(fyne.NewSize(800, 500))

	k := buildKeyboard(FULLKEYBOARDCOLUM)

	s := container.New(layout.NewHBoxLayout(), k)

	w.SetContent(s)

	w.ShowAndRun()
}

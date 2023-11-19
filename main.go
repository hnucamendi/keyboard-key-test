package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/hnucamendi/keyboard-key-test/generate"
)

const FULLKEYBOARDCOLUM int = 22

func main() {
	g := generate.Init()
	a := app.New()

	w := a.NewWindow("Keyboard Tester")
	w.Resize(fyne.NewSize(800, 500))

	k := g.GenerateKeyboard(FULLKEYBOARDCOLUM)

	w.SetContent(k)
	w.ShowAndRun()
}

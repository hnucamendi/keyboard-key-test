package generate

import (
	"fmt"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

type Generator struct {
	Nums       map[string]*canvas.Text
	Alphas     map[string]*canvas.Text
	Symbols    map[string]*canvas.Text
	AlphaNames map[string]*canvas.Text
	FNs        map[string]*canvas.Text

	channel chan string
}

var numKeySlice []string = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
var alphaKeySlice []string = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
var symbolsKeySlice []string = []string{"`", "~", "[", "{", "]", "}", "-", "=", "_", "+", "\\", "|", ";", ":", "'", "\"", ",", "<", ".", ">", "/", "?", "!", "@", "#", "$", "%", "^", "&", "*", "(", ")"}
var alphaNamedKeySlice []string = []string{"caps lock", "esc", "tab", "shift", "fn", "control", "option", "command", "alt", "ctrl", "delete", "prtsc", "scrlk", "pause", "ins", "home", "pgup", "del", "end", "pgdn", "num"}

func (g *Generator) generateNumKeys() {

	g.Nums = make(map[string]*canvas.Text)

	for _, nk := range numKeySlice {
		g.Nums[nk] = canvas.NewText(nk, color.White)
	}
}

func (g *Generator) generateAlphaKeys() {

	g.Alphas = make(map[string]*canvas.Text)

	for _, ak := range alphaKeySlice {
		g.Alphas[ak] = canvas.NewText(ak, color.White)
	}
}

func (g *Generator) generateSymbolKeys() {

	g.Symbols = make(map[string]*canvas.Text)

	for _, sk := range symbolsKeySlice {
		g.Symbols[sk] = canvas.NewText(sk, color.White)
	}
}

func (g *Generator) generateAlphaNamedKeys() {

	g.AlphaNames = make(map[string]*canvas.Text)

	for _, ank := range alphaNamedKeySlice {
		g.AlphaNames[ank] = canvas.NewText(ank, color.White)
	}
}

func (g *Generator) generateFnKeys() {

	g.FNs = make(map[string]*canvas.Text)

	fnSlice := []string{}
	for i := 0; i < 12; i++ {
		fnSlice = append(fnSlice, fmt.Sprintf("f%v", i+1))
	}

	for _, fn := range fnSlice {
		g.FNs[fn] = canvas.NewText(fn, color.White)
	}
}

func Init() *Generator {
	g := &Generator{}
	c := make(chan string)
	go func() {
		c <- "starting"
		g.generateAlphaNamedKeys()
		g.generateSymbolKeys()
		g.generateAlphaKeys()
		g.generateNumKeys()
		g.generateFnKeys()
	}()

	done := <-c

	fmt.Println(done)

	//msg := <-g.channel
	//fmt.Println(msg)

	return g
}

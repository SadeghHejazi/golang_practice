package main

import "github.com/eiannone/keyboard"

// listenKeyboard به‌صورت مداوم کلیدهای جهت‌دار را می‌خواند و جهت مار را عوض می‌کند
func listenKeyboard(g *Game) {
	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch key {
		case keyboard.KeyArrowUp:
			g.Direction = Point{X: 0, Y: -1}
		case keyboard.KeyArrowDown:
			g.Direction = Point{X: 0, Y: 1}
		case keyboard.KeyArrowLeft:
			g.Direction = Point{X: -1, Y: 0}
		case keyboard.KeyArrowRight:
			g.Direction = Point{X: 1, Y: 0}
		}
	}
}

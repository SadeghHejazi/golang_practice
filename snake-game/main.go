package main

import (
	"fmt"
	"time"

	"github.com/eiannone/keyboard"
)

func main() {
	fmt.Println("Welcome to SnakeGame")

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	game := Game{
		Snake: []Point{
			{X: 5, Y: 5},
			{X: 4, Y: 5},
			{X: 3, Y: 5},
		},
		Direction: Point{X: 1, Y: 0},
		Food:      Point{X: 8, Y: 5},
	}

	go listenKeyboard(&game)

	for {
		clearScreen()
		game.Move()
		game.Render()

		if game.GameOver {
			fmt.Println("Game Over!")
			break
		}

		time.Sleep(150 * time.Millisecond)
	}
}

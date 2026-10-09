package main

import "fmt"

// clearScreen صفحه‌ی ترمینال را پاک می‌کند
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// Render وضعیت فعلی بازی را روی ترمینال رسم می‌کند
func (g *Game) Render() {
	board := [height][width]string{}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			board[y][x] = "."
		}
	}

	for _, p := range g.Snake {
		board[p.Y][p.X] = "O"
	}

	board[g.Food.Y][g.Food.X] = "*"

	for _, row := range board {
		for _, cell := range row {
			fmt.Print(cell, " ")
		}
		fmt.Println()
	}
}

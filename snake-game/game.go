package main

import "math/rand"

const width = 35
const height = 10

type Game struct {
	Snake     []Point
	Direction Point
	Food      Point
	GameOver  bool
}

// Move یک قدم مار را جابه‌جا می‌کند، برخورد و خوردن غذا را بررسی می‌کند
func (g *Game) Move() {
	head := g.Snake[0]
	newHead := Point{
		X: head.X + g.Direction.X,
		Y: head.Y + g.Direction.Y,
	}

	for _, p := range g.Snake {
		if p == newHead {
			g.GameOver = true
			return
		}
	}

	g.Snake = append([]Point{newHead}, g.Snake...)

	if newHead == g.Food {
		g.SpawnFood()
	} else {
		g.Snake = g.Snake[:len(g.Snake)-1]
	}
}

// SpawnFood یک غذای جدید در نقطه‌ای خالی از صفحه قرار می‌دهد
func (g *Game) SpawnFood() {
	for {
		newFood := Point{
			X: rand.Intn(width),
			Y: rand.Intn(height),
		}

		occupied := false
		for _, p := range g.Snake {
			if p == newFood {
				occupied = true
				break
			}
		}

		if !occupied {
			g.Food = newFood
			return
		}
	}
}

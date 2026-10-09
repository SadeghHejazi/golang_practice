package main

import "fmt"

type Point struct {
	X, Y int
}

func IsValidWalk(walk []rune) bool {
	if len(walk) == 10 {
		X := 0
		Y := 0
		move := make(map[rune]Point, len(walk))
		move['n'] = Point{0, 1}
		move['s'] = Point{0, -1}
		move['w'] = Point{-1, 0}
		move['e'] = Point{+1, 0}
		for _, v := range walk {
			dark := move[v]
			Y += dark.Y
			X += dark.X
		}
		if X == 0 && Y == 0 {
			return true
		}
	}
	return false
}
func main() {
	fmt.Println("what is your input")
	var input string
	fmt.Scan(&input)
	fmt.Println(IsValidWalk([]rune(input)))
}

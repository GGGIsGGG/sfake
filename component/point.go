package component

import (
	"math/rand"
	"time"
)

type FoodPoint struct {
	Point
}

type Point struct {
	X int
	Y int
}

func (f *FoodPoint) RandomPos(playground *Playground, snake *Snake) *FoodPoint {
	x := getRandomAbove0(playground.Width - 3)
	y := getRandomAbove0(playground.Height - 3)
	if CheckInSnake(snake, x, y) {
		return f.RandomPos(playground, snake)
	}
	f.X = x
	f.Y = y
	return f
}

func CheckInSnake(s *Snake, x, y int) bool {
	for _, p := range s.Position {
		if p.X == x && p.Y == y {
			return true
		}
	}
	return false
}

func getRandomAbove0(n int) int {
	rand.Seed(time.Now().Unix())
	r := 0
	if r <= 0 {
		r = rand.Intn(n)
	}
	return r
}

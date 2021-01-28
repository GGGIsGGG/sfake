package component

import (
	"math/rand"
	"time"
)

const (
	SnakeBody              = 'S'
	Wall                   = 'W'
	Space                  = ' '
	Food                   = 'O'
	UP        DirectionCMD = 10 + iota
	DOWN
	LEFT
	RIGHT
)

type FoodPoint struct {
	Point
}

type Point struct {
	X int
	Y int
}

func (f *FoodPoint) RandomPos(playground *Playground, snake *Snake) *FoodPoint {
	x := getRandomAbove0(playground.Width - 1)
	y := getRandomAbove0(playground.Height - 1)
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
	for r <= 0 {
		r = rand.Intn(n)
	}
	return r
}

package component

import "time"

type Playground struct {
	Width  int
	Height int
	Ground [][]byte
}

func InitPlayground(rowCount, colCount int) *Playground {
	ground := make([][]byte, 0, colCount)
	playground := &Playground{
		rowCount,
		colCount,
		ground,
	}

	row := make([]byte, 0, playground.Width)
	temp := Wall
	for j := 1; j < playground.Width; j++ {
		temp += Wall
	}
	row = append(row, temp...)
	playground.Ground = append(playground.Ground, row)

	for i := 1; i < playground.Height-1; i++ {

		row = make([]byte, 0, playground.Width)
		temp := Wall
		for j := 1; j < playground.Width-1; j++ {
			temp += Space
		}
		temp += Wall
		row = append(row, temp...)
		playground.Ground = append(playground.Ground, row)
	}

	row = make([]byte, 0, playground.Width)
	temp = Wall
	for j := 1; j < playground.Width; j++ {
		temp += Wall
	}
	row = append(row, temp...)
	playground.Ground = append(playground.Ground, row)

	return playground
}

//CheckCrash 检查snake是否有撞墙
func (playground *Playground) CheckCrash(s *Snake) bool {
	for _, pos := range s.Position {
		if pos.X <= 0 || pos.Y <= 0 || pos.X >= playground.Width-1 || pos.Y >= playground.Height-1 {
			return true
		}
	}
	return false
}

func (playground *Playground) Reset() {

	for i := 0; i < playground.Width; i++ {
		playground.Ground[0][i] = []byte(Wall)[0]
	}

	for i := 1; i < playground.Height-1; i++ {
		playground.Ground[i][0] = []byte(Wall)[0]
		for j := 1; j < playground.Width-1; j++ {
			playground.Ground[i][j] = []byte(Space)[0]
		}
		playground.Ground[i][playground.Width-1] = []byte(Wall)[0]
	}

	for i := 0; i < playground.Width; i++ {
		playground.Ground[playground.Height-1][i] = []byte(Wall)[0]
	}
	time.Sleep(time.Second)
}

func (playground *Playground) DrawByPoint(p Point, b byte) {
	playground.Ground[p.Y][p.X] = b
}

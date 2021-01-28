package component

import "time"

type Playground struct {
	Width  int
	Height int
	Ground [][]rune
}

func InitPlayground(rowCount, colCount int) *Playground {
	ground := make([][]rune, 0, colCount)
	playground := &Playground{
		rowCount,
		colCount,
		ground,
	}

	row := make([]rune, 0, playground.Width)
	for j := 0; j < playground.Width; j++ {
		row = append(row, Wall)
	}
	playground.Ground = append(playground.Ground, row)

	for i := 1; i < playground.Height-1; i++ {

		row = make([]rune, 0, playground.Width)
		row = append(row, Wall)
		for j := 1; j < playground.Width-1; j++ {
			row = append(row, Space)
		}
		row = append(row, Wall)
		playground.Ground = append(playground.Ground, row)
	}

	row = make([]rune, 0, playground.Width)
	for j := 0; j < playground.Width; j++ {
		row = append(row, Space)
	}
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
		playground.Ground[0][i] = Wall
	}

	for i := 1; i < playground.Height-1; i++ {
		playground.Ground[i][0] = Wall
		for j := 1; j < playground.Width-1; j++ {
			playground.Ground[i][j] = Space
		}
		playground.Ground[i][playground.Width-1] = Wall
	}

	for i := 0; i < playground.Width; i++ {
		playground.Ground[playground.Height-1][i] = Wall
	}
	time.Sleep(time.Second)
}

func (playground *Playground) DrawByPoint(p Point, b rune) {
	playground.Ground[p.Y][p.X] = b
}

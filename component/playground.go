package component

type Playground struct {
	Width  int
	Height int
	Ground [][]byte
}

func InitPlayground(rowCount, colCount int) *Playground {
	ground := make([][]byte, 0, colCount)

	row := make([]byte, 0, rowCount)
	temp := Wall
	for j := 1; j < rowCount; j++ {
		temp += Wall
	}
	row = append(row, temp...)
	ground = append(ground, row)

	for i := 1; i < colCount-1; i++ {

		row = make([]byte, 0, rowCount)
		temp := Wall
		for j := 1; j < rowCount-1; j++ {
			temp += Space
		}
		temp += Wall
		row = append(row, temp...)
		ground = append(ground, row)
	}

	row = make([]byte, 0, rowCount)
	temp = Wall
	for j := 1; j < rowCount; j++ {
		temp += Wall
	}
	row = append(row, temp...)
	ground = append(ground, row)

	return &Playground{
		rowCount,
		colCount,
		ground,
	}
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

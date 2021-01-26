package component

type Playground struct {
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
		ground,
	}
}

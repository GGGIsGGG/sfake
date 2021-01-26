package component

type DirectionCMD int

const (
	SnakeBody              = "S"
	Wall                   = "W"
	Space                  = " "
	Food                   = "O"
	UP        DirectionCMD = 10 + iota
	DOWN
	LEFT
	RIGHT
)

type Snake struct {
	Direction DirectionCMD
	Position  []Point
}

func InitSnake(bodyLen int, playground *Playground) *Snake {
	ps := make([]Point, 0, bodyLen)

	originP := Point{
		len(playground.Ground[0]) / 2,
		len(playground.Ground) / 2,
	}

	ps = append(ps, originP)

	ox := originP.X
	oy := originP.Y
	for i := 0; i < bodyLen; i++ {

		nextP := Point{
			ox,
			oy,
		}
		oy++
		ps = append(ps, nextP)
	}

	return &Snake{
		UP,
		ps,
	}

}

//往前迈一步
func (s *Snake) Forward() {
	np := Point{}

	np.X = s.Position[s.snakeLen()-1].X
	np.Y = s.Position[s.snakeLen()-1].Y

	switch s.Direction {
	case UP:
		np.Y -= 1
	case DOWN:
		np.Y += 1
	case LEFT:
		np.X -= 1
	case RIGHT:
		np.X += 1
	}

	s.Position = append(s.Position, np)
}

//Shift 砍1格尾巴
func (s *Snake) Shift() {
	s.Position = s.Position[1:]
}

func (s *Snake) ChangeDirection(d DirectionCMD) {
	s.Direction = d
}

func (s *Snake) snakeLen() int {
	return len(s.Position)
}

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
	direction DirectionCMD
	Position  []Point
	//上一步所走的方向，用来检测下一步一定不能走相反方向。
	lastDirection DirectionCMD
}

func InitSnake(bodyLen int, playground *Playground) *Snake {
	ps := make([]Point, 0, bodyLen)

	originP := Point{
		playground.Width / 2,
		playground.Height / 2,
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
		DOWN,
		ps,
		DOWN,
	}

}

//Forward 往前迈一步
func (s *Snake) Forward() {
	np := Point{}

	np.X = s.Position[s.snakeLen()-1].X
	np.Y = s.Position[s.snakeLen()-1].Y

	switch s.direction {
	case UP:
		np.Y -= 1
	case DOWN:
		np.Y += 1
	case LEFT:
		np.X -= 1
	case RIGHT:
		np.X += 1
	}

	s.lastDirection = s.direction

	s.Position = append(s.Position, np)
}

//Shift 砍1格尾巴
func (s *Snake) Shift() Point {
	p := s.Position[0]
	s.Position = s.Position[1:]
	return p
}

func (s *Snake) ChangeDirection(d DirectionCMD) {
	switch d {
	case UP:
		if s.lastDirection != DOWN {
			s.direction = d
		}
	case DOWN:
		if s.lastDirection != UP {
			s.direction = d
		}
	case LEFT:
		if s.lastDirection != RIGHT {
			s.direction = d
		}
	case RIGHT:
		if s.lastDirection != LEFT {
			s.direction = d
		}
	}

}

func (s *Snake) snakeLen() int {
	return len(s.Position)
}

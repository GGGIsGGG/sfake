package controller

import (
	"sfake/component"
	"sfake/renderer"
)

type GameCore struct {
	snake      *component.Snake
	playground *component.Playground
	foodPoint  *component.FoodPoint
	scoreBoard *component.ScoreBoard
}

func NewGameCore() *GameCore {
	p := component.InitPlayground(40, 20)
	s := component.InitSnake(6, p)
	return &GameCore{
		s,
		p,
		(&component.FoodPoint{}).RandomPos(p, s),
		&component.ScoreBoard{},
	}
}

func (core *GameCore) Tick() {
	core.snake.Forward()

	if component.CheckInSnake(core.snake, core.foodPoint.X, core.foodPoint.Y) {
		core.foodPoint.RandomPos(core.playground, core.snake)
		core.scoreBoard.Score += 1
	} else {
		core.snake.Shift()
	}

	if core.playground.CheckCrash(core.snake) {
		core.GameOver(core.scoreBoard)
	} else {
		renderer.Render(core.playground, core.snake, core.foodPoint)
	}
}

func (core *GameCore) GameOver(board *component.ScoreBoard) {
	if board.Score > board.HighScore {
		board.HighScore = board.Score
	}
	renderer.ShowGameOver(board)
}

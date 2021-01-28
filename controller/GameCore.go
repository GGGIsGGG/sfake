package controller

import (
	"sfake/component"
	"sfake/renderer"
	"time"
)

type GameCore struct {
	speed      int64
	snake      *component.Snake
	playground *component.Playground
	foodPoint  *component.FoodPoint
	scoreBoard *component.ScoreBoard
	InputController
}

func NewGameCore(speed int64) *GameCore {
	p := component.InitPlayground(40, 20)
	s := component.InitSnake(6, p)
	gameCore := GameCore{
		int64(time.Second)/speed - 15*int64(time.Millisecond),
		s,
		p,
		(&component.FoodPoint{}).RandomPos(p, s),
		&component.ScoreBoard{},
		InputController{},
	}

	return &gameCore
}

func (core *GameCore) Tick() {
	core.snake.Forward()

	if component.CheckInSnake(core.snake, core.foodPoint.X, core.foodPoint.Y) {
		core.foodPoint.RandomPos(core.playground, core.snake)
		core.scoreBoard.Score += 1
	} else {
		core.Shift()
	}

	if core.playground.CheckCrash(core.snake) || core.snake.CheckSelfCrash() {
		core.GameOver()
	} else {
		renderer.Render(core.playground, core.snake, core.foodPoint)
	}
}

func (core *GameCore) Shift() {
	p := core.snake.Shift()
	core.playground.DrawByPoint(p, component.Space)
}

func (core *GameCore) GameOver() {
	core.stop = true
	if core.scoreBoard.Score > core.scoreBoard.HighScore {
		core.scoreBoard.HighScore = core.scoreBoard.Score
	}
	renderer.ShowGameOver(core.scoreBoard)
	core.Reset()
	core.RegisterMenuDownEvent()
	core.GameStart(1)
}

func (core *GameCore) GameStart(rate int64) {
	core.stop = false
	core.RegisterSnakeKeyDownEvent(core.snake)
	for !core.stop {
		core.Tick()
		time.Sleep(time.Duration(core.speed))
	}
}

func (core *GameCore) Reset() {
	core.scoreBoard.Score = 0
	core.playground.Reset()
	core.snake = component.InitSnake(6, core.playground)
	core.foodPoint.RandomPos(core.playground, core.snake)
}

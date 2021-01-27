package controller

import (
	"sfake/component"
	"sfake/renderer"
)

type GameController struct {
	snake      *component.Snake
	playground *component.Playground
	scoreBoard *component.ScoreBoard
	foodPoint  *component.FoodPoint
}

func NewGameController() *GameController {
	p := component.InitPlayground(40, 20)
	s := component.InitSnake(6, p)
	return &GameController{
		s,
		p,
		&component.ScoreBoard{},
		(&component.FoodPoint{}).RandomPos(p, s),
	}
}

func (ctrl *GameController) Tick() {
	renderer.Render(ctrl.playground, ctrl.snake, ctrl.foodPoint)
}

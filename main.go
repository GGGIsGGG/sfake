package main

import (
	"sfake/controller"
	"sfake/renderer"
)

func main() {
	renderer.Clear()
	gameCore := controller.NewGameCore(4)
	//gameCore.GameStart(1)
	gameCore.GameOver()
}

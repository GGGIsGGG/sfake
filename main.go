package main

import (
	"sfake/controller"
	"sfake/renderer"
)

func main() {
	renderer.Clear()
	gameCore := controller.NewGameController()
	gameCore.Tick()
}

package main

import (
	"sfake/component"
	"sfake/renderer"
)

func main() {
	renderer.Clear()
	p := component.InitPlayground(40, 20)
	s := component.InitSnake(6, p)

	renderer.Render(p, s)
}

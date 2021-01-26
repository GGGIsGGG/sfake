package renderer

import (
	"fmt"
	"os"
	"os/exec"
	"sfake/component"
)

func Render(playground *component.Playground, snake *component.Snake) {

	view := getView(playground.Ground, snake)
	for _, raw := range view {
		fmt.Println(string(raw))
	}
}

func getView(bs [][]byte, snake *component.Snake) [][]byte {
	for _, body := range snake.Position {
		bs[body.Y][body.X] = []byte(component.SnakeBody)[0]
	}
	return bs
}

func Clear() {
	cmd := exec.Command("cmd.exe", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

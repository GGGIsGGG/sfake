package renderer

import (
	"fmt"
	"os"
	"os/exec"
	"sfake/component"
)

func Render(playground *component.Playground, snake *component.Snake, point *component.FoodPoint) {
	Clear()
	view := getView(playground.Ground, snake, point)
	for _, raw := range view {
		fmt.Println(string(raw))
	}
}

func getView(bs [][]byte, snake *component.Snake, point *component.FoodPoint) [][]byte {
	bs[point.Y][point.X] = []byte(component.Food)[0]
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

func GameOver(board *component.ScoreBoard) {
	Clear()
	fmt.Println("Game Over!")
	fmt.Printf("High Score: %d", board.HighScore)
	fmt.Printf("Score: %d", board.Score)
}

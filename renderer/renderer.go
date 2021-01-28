package renderer

import (
	"fmt"
	"os"
	"os/exec"
	"sfake/component"
)

//var cost int64

func Render(playground *component.Playground, snake *component.Snake, point *component.FoodPoint) {

	Clear()
	view := getView(playground.Ground, snake, point)
	for _, row := range view {
		fmt.Println(string(row))
	}
	//ok := time.Duration(time.Now().UnixNano() - cost).Milliseconds()
	//cost = time.Now().UnixNano()
	//fmt.Printf("render cost: %d", ok)
}

func getView(bs [][]rune, snake *component.Snake, point *component.FoodPoint) [][]rune {
	bs[point.Y][point.X] = component.Food
	for _, body := range snake.Position {
		bs[body.Y][body.X] = component.SnakeBody
	}
	return bs
}

func Clear() {
	cmd := exec.Command("cmd.exe", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ShowGameOver(board *component.ScoreBoard) {
	Clear()
	fmt.Println("Game Over!")
	fmt.Printf("High Score: %d\n", board.HighScore)
	fmt.Printf("Score: %d\n", board.Score)
	fmt.Printf("press Space to start\n")
}

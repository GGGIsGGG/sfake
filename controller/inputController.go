package controller

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func GetScan() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter your name:")
	go shit(inputReader, 'f')
	//switch input {
	//case "yinzhengjie\n":
	//	fmt.Println("Welcome yinzhengjie!")
	//case "bingan\n":
	//	fmt.Println("Welcome bingan!")
	//case "liufei\n":
	//	fmt.Println("Welcome liufei")
	//default:
	//	fmt.Println("You are not welcome here! Goodbye!")
	//}
	/*    //version 2:
	      switch input {
	      case "yinzhengjie\n":
	          fallthrough
	      case "jiashanpeng\n":
	          fallthrough
	      case "hansenyu\n":
	          fmt.Printf("Welcome %s\n", input)
	      default:
	          fmt.Printf("You are not welcome here! Goodbye!\n")
	      }

	      // version 3:
	      switch input {
	      case "yinzhengjie\n", "wuzhiguang\n":
	          fmt.Printf("Welcome %s", input)
	      default:
	          fmt.Printf("You are not welcome here! Goodbye!\n")
	      }

	*/
	time.Sleep(time.Second * 20)
}

func shit(inputReader *bufio.Reader, n byte) {
	input, err := inputReader.ReadString(n)

	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	fmt.Printf("Your name is %s", input)
}

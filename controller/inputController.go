package controller

import (
	"github.com/nsf/termbox-go"
	"sfake/component"
)

func init() {
	err := termbox.Init()
	if err != nil {

		panic(err)

	}
}

type InputController struct {
	stop bool
}

func (ic *InputController) RegisterMenuDownEvent() {
loop:
	for ic.stop {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeySpace:
				break loop
			}
		}
	}
}

func (ic *InputController) RegisterSnakeKeyDownEvent(s *component.Snake) {
	go func() {
		for !ic.stop {
			switch ev := termbox.PollEvent(); ev.Type {

			case termbox.EventKey:
				if ev.Ch != 0 {
					switch ev.Ch {
					case 'w':
						s.ChangeDirection(component.UP)
					case 's':
						s.ChangeDirection(component.DOWN)
					case 'a':
						s.ChangeDirection(component.LEFT)
					case 'd':
						s.ChangeDirection(component.RIGHT)
					}
				} else {
					switch ev.Key {

					case termbox.KeyArrowUp:
						s.ChangeDirection(component.UP)
					case termbox.KeyArrowDown:
						s.ChangeDirection(component.DOWN)
					case termbox.KeyArrowLeft:
						s.ChangeDirection(component.LEFT)
					case termbox.KeyArrowRight:
						s.ChangeDirection(component.RIGHT)

					default:

					}
				}
			}
		}
	}()
}

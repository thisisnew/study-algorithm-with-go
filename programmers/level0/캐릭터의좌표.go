package main

import "fmt"

func main() {
	fmt.Println(캐릭터의좌표([]string{"down", "down", "down", "down", "down"}, []int{7, 9}))
}

func 캐릭터의좌표(keyinput []string, board []int) []int {

	var x = 0
	var y = 0
	xMax := (board[0] - 1) / 2
	yMax := (board[1] - 1) / 2

	for _, k := range keyinput {

		switch k {
		case "left":
			temp := x - 1

			if temp >= xMax*-1 {
				x = temp
			}

		case "right":
			temp := x + 1

			if temp <= xMax {
				x = temp
			}
		case "up":
			temp := y + 1

			if temp <= yMax {
				y = temp
			}
		case "down":
			temp := y - 1

			if temp >= yMax*-1 {
				y = temp
			}
		}

	}

	return []int{x, y}
}

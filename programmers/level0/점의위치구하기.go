package main

import "fmt"

func main() {
	fmt.Println(점의위치구하기([]int{2, 4}))
}

func 점의위치구하기(dot []int) int {

	x := dot[0]
	y := dot[1]

	switch {
	case x < 0 && y > 0:
		return 2
	case x < 0 && y < 0:
		return 3
	case x > 0 && y < 0:
		return 4

	}

	return 1
}

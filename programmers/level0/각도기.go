package main

import "fmt"

func main() {
	fmt.Println(각도기(70))
}

func 각도기(angle int) int {

	switch {
	case angle > 0 && angle < 90:
		return 1
	case angle == 90:
		return 2
	case angle > 90 && angle < 180:
		return 3
	case angle == 180:
		return 4
	}

	return 0
}

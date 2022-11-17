package main

import "fmt"

func main() {
	fmt.Println(개미군단(23))
}

func 개미군단(hp int) int {

	var result int

	for {
		if hp == 0 {
			return result
		}

		switch {
		case hp >= 5:
			hp -= 5
		case hp >= 3:
			hp -= 3
		default:
			hp -= 1
		}

		result++
	}

}

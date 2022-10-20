package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(예상대진표(8, 4, 7))
}

func 예상대진표(n int, a int, b int) int {

	var result int

	for n > 1 {
		result++

		if math.Abs(float64(a-b)) == 1 {
			return result
		}

		if a%2 == 0 {
			a = a / 2
		} else {
			a = a/2 + 1
		}

		if b%2 == 0 {
			b = b / 2
		} else {
			b = b/2 + 1
		}

		n = n / 2
	}

	return result
}

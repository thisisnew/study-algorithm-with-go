package main

import (
	"fmt"
)

func main() {
	fmt.Println(하샤드수(10))
}

func 하샤드수(x int) bool {
	return x%sumDigit(x) == 0
}

func sumDigit(x int) int {
	var result int

	for {
		if x == 0 {
			break
		}

		result += x % 10
		x = x / 10
	}

	return result
}

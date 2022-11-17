package main

import (
	"fmt"
)

func main() {
	fmt.Println(세균증식(2, 10))
}

func 세균증식(n int, t int) int {

	var result = n

	for i := 1; i <= t; i++ {
		result *= 2
	}

	return result
}

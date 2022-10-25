package main

import (
	"fmt"
)

func main() {
	fmt.Println(콜라문제(3, 1, 20))
}

func 콜라문제(a int, b int, n int) int {

	var empty = n
	var result = 0

	for {
		e := empty / a * b
		r := empty % a
		result += e
		empty = e + r

		if empty < a {
			break
		}
	}

	return result
}

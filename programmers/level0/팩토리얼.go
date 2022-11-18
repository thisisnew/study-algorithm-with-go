package main

import "fmt"

func main() {
	fmt.Println(팩토리얼(7))
}

func 팩토리얼(n int) int {

	var idx = 1

	for {
		if fact(idx) > n {
			return idx - 1
		}

		idx++
	}

}

func fact(n int) int {

	if n <= 1 {
		return n
	}

	return n * fact(n-1)
}

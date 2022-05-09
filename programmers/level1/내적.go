package main

import "fmt"

func main() {
	a := []int{-1, 0, 1}
	b := []int{1, 0, -1}

	fmt.Println(내적(a, b))
}

func 내적(a []int, b []int) int {

	var result int

	for i, el := range a {
		result += el * b[i]
	}

	return result
}

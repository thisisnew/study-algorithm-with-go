package main

import "fmt"

func main() {
	fmt.Println(
		할인행사(
			[]string{"banana", "apple", "rice", "pork", "pot"},
			[]int{3, 2, 2, 2, 1},
			[]string{"chicken", "apple", "apple", "banana", "rice", "apple", "pork", "banana", "pork", "rice", "pot", "banana", "apple", "banana"},
		),
	)
}

func 할인행사(want []string, number []int, discount []string) int {
	return 0
}

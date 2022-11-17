package main

import "fmt"

func main() {
	fmt.Println(중복된숫자개수([]int{1, 1, 2, 3, 4, 5}, 1))
}

func 중복된숫자개수(array []int, n int) int {

	var result int

	for _, num := range array {
		if num == n {
			result++
		}
	}

	return result
}

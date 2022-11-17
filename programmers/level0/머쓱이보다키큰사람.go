package main

import "fmt"

func main() {
	fmt.Println([]int{149, 180, 192, 170}, 167)
}

func 머쓱이보다키큰사람(array []int, height int) int {

	var result int

	for _, h := range array {
		if h > height {
			result++
		}
	}

	return result
}

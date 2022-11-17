package main

import "fmt"

func main() {
	fmt.Println(가장큰수찾기([]int{1, 8, 3}))
}

func 가장큰수찾기(array []int) []int {

	var max = 0
	var idx = 0

	for i, n := range array {

		if n > max {
			max = n
			idx = i
		}

	}

	return []int{max, idx}
}

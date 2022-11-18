package main

import "fmt"

func main() {
	fmt.Println(공던지기([]int{1, 2, 3, 4, 5, 6}, 5))
}

func 공던지기(numbers []int, k int) int {

	var idx = 0
	var cnt = 1
	for {
		if cnt == k {
			return numbers[idx]
		}

		cnt++
		idx += 2
		if idx >= len(numbers) {
			idx -= len(numbers)
		}
	}
}

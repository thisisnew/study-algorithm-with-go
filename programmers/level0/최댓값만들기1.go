package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(최댓값만들기1([]int{0, 31, 24, 10, 1, 9}))
}

func 최댓값만들기1(numbers []int) int {

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] > numbers[j]
	})

	return numbers[0] * numbers[1]
}

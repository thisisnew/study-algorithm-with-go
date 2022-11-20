package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(최댓값만들기2([]int{-1, 1}))
}

func 최댓값만들기2(numbers []int) int {

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	var result = numbers[0] * numbers[1]

	if numbers[len(numbers)-1]*numbers[len(numbers)-2] > result {
		result = numbers[len(numbers)-1] * numbers[len(numbers)-2]
	}

	return result
}

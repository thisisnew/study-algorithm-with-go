package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(중앙값구하기([]int{1, 2, 7, 10, 11}))
}

func 중앙값구하기(array []int) int {

	sort.Slice(array, func(i, j int) bool {
		return array[i] < array[j]
	})

	return array[(len(array)-1)/2]
}

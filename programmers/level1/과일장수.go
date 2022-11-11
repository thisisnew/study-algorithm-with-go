package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(과일장수(4, 3, []int{4, 1, 2, 2, 4, 4, 4, 4, 1, 2, 4, 2}))
}

func 과일장수(k int, m int, score []int) int {
	return calMaximumProfit(loadApples(m, score))
}

func loadApples(m int, score []int) [][]int {
	sort.Slice(score, func(i, j int) bool {
		return score[i] > score[j]
	})

	var appleBoxes [][]int
	var apples = make([]int, m)
	var idx = 0

	for i := 0; i < len(score); i++ {
		apples[idx] = score[i]
		idx++

		if idx == m {
			idx = 0
			appleBoxes = append(appleBoxes, apples)
			apples = make([]int, m)
		}
	}

	return appleBoxes
}

func calMaximumProfit(appleBoxes [][]int) int {

	var result = 0
	var boxes = len(appleBoxes)

	for _, apples := range appleBoxes {
		m := len(apples)
		result += apples[m-1] * m * boxes
	}

	return result
}

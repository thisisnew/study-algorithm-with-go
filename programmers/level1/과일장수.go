package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(과일장수(4, 3, []int{4, 1, 2, 2, 4, 4, 4, 4, 1, 2, 4, 2}))
}

func 과일장수(k int, m int, score []int) int {
	return calMaximumProfit(countAppleBoxes(m, loadApples(m, score)))
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

func countAppleBoxes(m int, appleBoxes [][]int) map[int]int {

	var applesMap = map[int]int{}

	for _, apples := range appleBoxes {
		minPrice := apples[len(apples)-1]
		price := m * minPrice
		applesMap[price]++
	}

	return applesMap
}

func calMaximumProfit(applesMap map[int]int) int {

	var result = 0

	for price, boxes := range applesMap {
		result += price * boxes
	}

	return result
}

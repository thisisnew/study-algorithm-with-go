package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(과일장수(4, 3, []int{4, 1, 2, 2, 4, 4, 4, 4, 1, 2, 4, 2}))
}

func 과일장수(k int, m int, score []int) int {
	return calMaximumProfit(m, countAppleBoxes(loadApples(m, score)))
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

func countAppleBoxes(appleBoxes [][]int) [][]int {

	if len(appleBoxes) == 0 {
		return [][]int{}
	}

	var a = appleBoxes[0]
	var boxes = 1
	var result [][]int

	for i := 1; i < len(appleBoxes); i++ {
		apples := appleBoxes[i]

		if !isSameAppleBox(a, apples) {
			result = addBoxCntToAppleBox(result, a, boxes)
			a = apples
			boxes = 0
		}

		boxes++
	}

	result = addBoxCntToAppleBox(result, a, boxes)

	return result
}

func isSameAppleBox(ab, ac []int) bool {

	for i, a := range ab {
		if a != ac[i] {
			return false
		}
	}

	return true
}

func addBoxCntToAppleBox(appleBoxes [][]int, apples []int, boxes int) [][]int {

	apples = append(apples, boxes)
	appleBoxes = append(appleBoxes, apples)

	return appleBoxes
}

func calMaximumProfit(m int, appleBoxes [][]int) int {

	var result = 0

	for _, apples := range appleBoxes {
		ln := len(apples)
		minPrice := apples[ln-2]
		boxes := apples[ln-1]
		result += minPrice * m * boxes
	}

	return result
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(과일장수(4, 3, []int{4, 1, 2, 2, 4, 4, 4, 4, 1, 2, 4, 2}))
}

func 과일장수(k int, m int, score []int) int {
	return calMaximumProfit(countAppleBoxes(loadApples(m, score)))
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

	var a = appleBoxes[0]
	var boxes = 1

	for i := 1; i < len(appleBoxes); i++ {
		apples := appleBoxes[i]

		if !isSameAppleBox(a, apples) || i == len(appleBoxes)-1 {
			addBoxCntToAppleBox(&appleBoxes, i, boxes)
			a = apples
			boxes = 0
		}

		boxes++
	}

	return appleBoxes
}

func isSameAppleBox(ab, ac []int) bool {

	for i, a := range ab {
		if a != ac[i] {
			return false
		}
	}

	return true
}

func addBoxCntToAppleBox(appleBoxes *[][]int, idx, boxes int) {

	for i := idx - boxes; i < idx; i++ {
		apples := (*appleBoxes)[i]
		apples = append(apples, boxes)
		(*appleBoxes)[i] = apples
	}

}

func calMaximumProfit(appleBoxes [][]int) int {

	var result = 0

	for _, apples := range appleBoxes {
		m := len(apples)
		boxes := apples[m-1]
		minimum := apples[m-2]
		result += minimum * m * boxes
	}

	return result
}

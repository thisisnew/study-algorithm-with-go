package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(과일장수(3, 4, []int{1, 2, 3, 1, 2, 3, 1}))
}

func 과일장수(k int, m int, score []int) int {

	sort.Slice(score, func(i, j int) bool {
		return score[i] > score[j]
	})

	var boxCnt = 0
	var minimums []int
	for i := 0; i < len(score)-m; i += m {
		boxCnt++

		minimums = append(minimums, score[i+m-1])
	}

	return calMaximumProfit(boxCnt, minimums)
}

func calMaximumProfit(boxCnt int, minimums []int) int {

	var result = 0

	for _, m := range minimums {
		result += m * boxCnt
	}

	return result
}

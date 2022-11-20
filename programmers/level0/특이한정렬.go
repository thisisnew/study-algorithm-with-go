package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(특이한정렬([]int{1, 2, 3, 4, 5, 6}, 4))
}

func 특이한정렬(numlist []int, n int) []int {

	sort.Slice(numlist, func(i, j int) bool {

		var difI = getDifBetweenScores(n, numlist[i])
		var difJ = getDifBetweenScores(n, numlist[j])

		if difI == difJ {
			return numlist[i] > numlist[j]
		}

		return difI < difJ
	})

	return numlist
}

func getDifBetweenScores(s1, s2 int) int {

	if s1 > s2 {
		return s1 - s2
	}

	return s2 - s1

}

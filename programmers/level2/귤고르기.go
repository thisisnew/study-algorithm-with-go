package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(귤고르기(2, []int{1, 1, 1, 1, 2, 2, 2, 3}))
}

func 귤고르기(k int, tangerine []int) int {

	var counts = countTangerines(getTangerineBasket(tangerine))
	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})

	return countMinTangerineTypes(k, counts)
}

func getTangerineBasket(tangerine []int) map[int]int {
	var result = map[int]int{}

	for _, t := range tangerine {
		result[t]++
	}

	return result
}

func countTangerines(tangerineBasket map[int]int) []int {
	var result []int

	for _, v := range tangerineBasket {
		result = append(result, v)
	}

	return result
}

func countMinTangerineTypes(k int, counts []int) int {

	var result = 0
	var cnt = 0

	for i := 0; i < len(counts); i++ {
		if cnt >= k {
			return result
		}

		result++
		cnt += counts[i]
	}

	return result
}

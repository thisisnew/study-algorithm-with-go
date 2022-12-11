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

	_, result := countFittedMinTangerineTypes(k, counts)

	return result
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

func countFittedMinTangerineTypes(k int, counts []int) (int, int) {

	var result = 0
	var cnt = 0

	for i := 0; i < len(counts); i++ {
		cnt = counts[i]
		result = 0

		for j := i + 1; j < len(counts); j++ {
			result++

			if cnt == k {
				return cnt, result
			}

			if cnt+counts[j] > k {
				continue
			}

			cnt += counts[j]
		}
	}

	return cnt, result
}

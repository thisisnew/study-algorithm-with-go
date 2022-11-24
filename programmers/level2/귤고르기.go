package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(귤고르기(2, []int{1, 1, 1, 1, 2, 2, 2, 3}))
}

func 귤고르기(k int, tangerine []int) int {

	var tangerineBasket = getTangerineBasket(tangerine)
	var counts = countTangerines(tangerineBasket)
	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})

	return getMinCountTangerines(k, counts) + 1
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

func getMinCountTangerines(k int, counts []int) int {
	var idx = 0

	for i := 0; i < len(counts); i++ {
		idx = 0
		var sum = 0

		for j := i; j < len(counts); j++ {
			c := counts[j]

			if sum+c > k {
				continue
			}

			sum += c

			if sum == k {
				idx = j
				return idx
			}
		}
	}

	return idx
}

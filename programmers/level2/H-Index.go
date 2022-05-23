package main

import (
	"fmt"
	"sort"
)

func main() {

	citations := []int{3, 0, 6, 1, 5}

	fmt.Println(hIndex(citations))

}

func hIndex(citations []int) int {

	var result, a, b int

	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	for {

		if citations[0] <= result {
			result = len(citations)
			break
		}

		a = above(result, citations)
		b = below(result, citations)

		if a != result || b != result {
			result++
			continue
		}

		break
	}

	return result
}

func above(n int, citations []int) int {

	var result = 0

	for _, c := range citations {

		if c < n {
			continue
		}

		result++

		if result > n {
			break
		}
	}

	return result

}

func below(n int, citations []int) int {

	var result = 0

	for _, c := range citations {

		if c > n {
			continue
		}

		result++

		if result > n {
			break
		}
	}

	return result

}

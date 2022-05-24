package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {

	citations := []int{3, 0, 6, 1, 5}

	fmt.Println(hIndex(citations))

}

func hIndex(citations []int) int {

	var result int

	sort.Slice(citations, func(i, j int) bool {
		return citations[i] < citations[j]
	})

	lc := len(citations)

	for i, c := range citations {

		s := math.Min(float64(c), float64(lc-i))

		if int(s) >= result {
			result = int(math.Max(float64(result), s))
			continue
		}

		break
	}

	return result
}

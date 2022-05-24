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

	var result int

	sort.Slice(citations, func(i, j int) bool {
		return citations[i] < citations[j]
	})

	for i, c := range citations {
		if c < len(citations)-i {
			continue
		}

		result = len(citations) - i
		break
	}

	return result
}

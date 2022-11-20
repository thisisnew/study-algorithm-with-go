package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(평행([][]int{{1, 4}, {9, 2}, {3, 8}, {11, 6}}))
}

func 평행(dots [][]int) int {

	sort.Slice(dots, func(i, j int) bool {
		return dots[i][0] < dots[j][0]
	})

	return 0
}

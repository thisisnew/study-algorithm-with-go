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

	dot1 := dots[0]
	dot2 := dots[1]
	dot3 := dots[2]
	dot4 := dots[3]

	if getInclinationDots(dot1, dot2, dot3, dot4) {

	}

	if getInclinationDots(dot1, dot3, dot2, dot4) {

	}

	if getInclinationDots(dot1, dot4, dot3, dot4) {

	}

	return 0
}

func getInclinationDots(dot1, dot2, dot3, dot4 []int) bool {
	return (dot2[1]-dot1[1])/(dot2[0]-dot1[0]) == (dot4[1]-dot3[1])/(dot4[0]-dot3[0])
}

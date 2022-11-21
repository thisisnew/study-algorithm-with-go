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
		if (dot4[1]-dot3[1])/(dot2[1]-dot1[1]) != -1 {
			return 1
		}
	}

	if getInclinationDots(dot1, dot3, dot2, dot4) {
		if (dot4[1]-dot2[1])/(dot3[1]-dot1[1]) != -1 {
			return 1
		}
	}

	if getInclinationDots(dot1, dot4, dot2, dot3) {
		if (dot3[1]-dot2[1])/(dot4[1]-dot1[1]) != -1 {
			return 1
		}
	}

	return 0
}

func getInclinationDots(dot1, dot2, dot3, dot4 []int) bool {

	if dot2[0]-dot1[0] == 0 && dot4[0]-dot3[0] == 0 {
		if dot2[1]-dot1[1] == 0 && dot4[1]-dot3[1] == 0 {
			return false
		}
		return true
	}

	if dot2[0]-dot1[0] != 0 && dot4[0]-dot3[0] == 0 {
		return (dot2[1]-dot1[1])/(dot2[0]-dot1[0]) == 1
	}

	if dot2[0]-dot1[0] == 0 && dot4[0]-dot3[0] != 0 {
		return (dot4[1]-dot3[1])/(dot4[0]-dot3[0]) == 1
	}

	return (dot2[1]-dot1[1])/(dot2[0]-dot1[0]) == (dot4[1]-dot3[1])/(dot4[0]-dot3[0])
}

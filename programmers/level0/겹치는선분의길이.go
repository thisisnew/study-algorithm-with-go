package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(겹치는선분의길이([][]int{{0, 5}, {3, 9}, {1, 10}}))
}

func 겹치는선분의길이(lines [][]int) int {

	sort.Slice(lines, func(i, j int) bool {
		return lines[i][0] < lines[j][0]
	})

	var result = make([]int, 2)

	for i := 0; i < len(lines)-1; i++ {

		lineI := lines[i]

		for j := i + 1; j < len(lines); j++ {
			lineJ := lines[j]

			if lineI[0] > lineJ[0] {
				result[0] = lineI[0]
			} else {
				result[0] = lineJ[0]
			}

			if lineI[1] > lineJ[1] {
				result[1] = lineJ[1]
			} else {
				result[1] = lineI[1]
			}
		}
	}

	return result[1] - result[0]
}

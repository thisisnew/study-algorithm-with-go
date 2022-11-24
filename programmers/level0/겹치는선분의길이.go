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

	var dup [][]int

	for i := 0; i < len(lines)-1; i++ {
		var line = make([]int, 2)
		lineI := lines[i]

		for j := i + 1; j < len(lines); j++ {
			lineJ := lines[j]

			if lineI[0] > lineJ[0] {
				line[0] = lineI[0]
			} else {
				line[0] = lineJ[0]
			}

			if lineJ[1] < lineI[1] {
				line[1] = lineJ[1]
			} else {
				line[1] = lineI[1]
			}
		}

		if line[0] != 0 && line[1] != 0 {
			dup = append(dup, line)
		}
	}

	return 0
}

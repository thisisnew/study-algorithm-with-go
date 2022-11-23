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

	line1 := lines[0]
	line2 := lines[1]
	line3 := lines[2]

	if line1[0] <= line2[0] && line1[0] <= line3[0] && line1[1] >= line2[1] && line1[1] >= line3[1] {
		return line1[1] - line1[0]
	}

	if line2[0] <= line1[0] && line2[0] <= line3[0] && line2[1] >= line1[1] && line2[1] >= line3[1] {
		return line2[1] - line2[0]
	}

	if line3[0] <= line1[0] && line3[0] <= line2[0] && line3[1] >= line2[1] && line3[1] >= line1[1] {
		return line3[1] - line3[0]
	}

	if isContainLine(line1, line2, line3) {
		line1 = nil
	}

	if isContainLine(line2, line1, line3) {
		line2 = nil
	}

	if isContainLine(line3, line1, line2) {
		line3 = nil
	}

	var result int

	if line1 != nil && line2 != nil && line1[1] > line2[0] {
		if line1[1] > line2[1] {
			if line3 != nil && line1[1] > line3[0] {
				if line1[1] > line3[1] {
					result += (line2[1] - line2[0]) + (line3[1] - line3[0])
				} else {
					result += (line2[1] - line2[0]) + (line1[1] - line3[0])
				}
			} else {
				result += line2[1] - line2[0]
			}
		} else {
			result += line1[1] - line2[0]
		}
	}

	if line2 != nil && line3 != nil && line2[1] > line3[0] {
		if line2[1] > line3[1] {
			result += line3[1] - line3[0]
		} else {
			result += line2[1] - line3[0]
		}
	}

	return result
}

func isContainLine(line1, line2, line3 []int) bool {
	return (line1[0] <= line2[0] && line1[1] >= line2[1]) || (line1[0] <= line3[0] && line1[1] >= line3[1])
}

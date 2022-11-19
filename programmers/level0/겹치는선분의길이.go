package main

import (
	"fmt"
)

func main() {
	fmt.Println(겹치는선분의길이([][]int{{0, 5}, {3, 9}, {1, 10}}))
}

func 겹치는선분의길이(lines [][]int) int {

	line1 := lines[0]
	line2 := lines[1]
	line3 := lines[2]

	var result int

	if line1[1] > line2[0] {
		if line1[1] > line2[1] {
			if line1[1] > line3[0] {
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

	if line2[1] > line3[0] {
		if line2[1] > line3[1] {
			result += line3[1] - line3[0]
		} else {
			result += line2[1] - line3[0]
		}
	}

	return result
}

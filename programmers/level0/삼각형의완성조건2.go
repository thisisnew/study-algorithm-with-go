package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(삼각형의완성조건2([]int{3, 6}))
}

func 삼각형의완성조건2(sides []int) int {

	sort.Slice(sides, func(i, j int) bool {
		return sides[i] > sides[j]
	})

	max := sides[0]

	ln := 0
	var result = 0

	for {
		ln++

		if ln <= max {
			if max >= ln+sides[1] {
				continue
			}

			result++
		} else {
			if ln >= max+sides[1] {
				break
			}

			result++
		}
	}

	return result
}

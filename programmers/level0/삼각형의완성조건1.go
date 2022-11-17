package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(삼각형의완성조건1([]int{1, 2, 3}))
}

func 삼각형의완성조건1(sides []int) int {

	sort.Slice(sides, func(i, j int) bool {
		return sides[i] < sides[j]
	})

	if sides[0]+sides[1] <= sides[2] {
		return 2
	}

	return 1
}

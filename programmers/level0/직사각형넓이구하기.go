package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(직사각형넓이구하기([][]int{{1, 1}, {2, 1}, {2, 2}, {1, 2}}))
}

func 직사각형넓이구하기(dots [][]int) int {

	var w = make([]int, 4)
	var h = make([]int, 4)

	for i, dot := range dots {
		w[i] = dot[0]
		h[i] = dot[1]
	}

	sort.Slice(w, func(i, j int) bool {
		return w[i] < w[j]
	})

	sort.Slice(h, func(i, j int) bool {
		return h[i] < h[j]
	})

	return (w[3] - w[0]) * (h[3] - h[0])
}

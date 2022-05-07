package main

import (
	"fmt"
)

func main() {

	n := []int{
		4, 3, 1, 2,
	}
	fmt.Println(제일작은수제거하기(n))
}

func 제일작은수제거하기(arr []int) []int {

	ln := len(arr)

	if ln <= 1 {
		return []int{-1}
	}

	var minIdx = 0
	var minValue = arr[0]

	for i, v := range arr {

		if v >= minValue {
			continue
		}

		minValue = v
		minIdx = i
	}

	if minIdx == 0 {
		return arr[1:]
	}

	if minIdx == ln-1 {
		return arr[:minIdx]
	}

	result := arr[:minIdx]
	result = append(result, arr[minIdx+1:]...)

	return result
}

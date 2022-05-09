package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{3, 2, 6}
	divisor := 10
	fmt.Println(나누어떨어지는숫자배열(arr, divisor))
}

func 나누어떨어지는숫자배열(arr []int, divisor int) []int {

	var result []int

	for _, prop := range arr {

		if prop%divisor != 0 {
			continue
		}

		result = append(result, prop)
	}

	if len(result) == 0 {
		return []int{-1}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return result
}

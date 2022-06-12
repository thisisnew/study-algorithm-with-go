package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(N개의최소공배수([]int{1, 2, 3}))
}

func N개의최소공배수(arr []int) int {

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

	var result = arr[0]

	for {
		if !isLcm(arr, result) {
			result++
			continue
		}

		return result
	}

}

func isLcm(arr []int, lcm int) bool {

	for _, n := range arr {
		if lcm%n != 0 {
			return false
		}
	}

	return true
}

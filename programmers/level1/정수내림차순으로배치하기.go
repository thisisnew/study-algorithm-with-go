package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(정수내림차순으로배치하기(118372))
}

func 정수내림차순으로배치하기(n int64) int64 {

	var numSlice []int64

	for {

		if n == 0 {
			break
		}

		numSlice = append(numSlice, n%10)

		n = n / 10

	}

	sort.Slice(numSlice, func(i, j int) bool {
		return numSlice[i] > numSlice[j]
	})

	var result int64

	for _, num := range numSlice {

		result *= 10
		result += num

	}

	return result
}

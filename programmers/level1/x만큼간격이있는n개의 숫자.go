package main

import "fmt"

func main() {
	fmt.Println(x만큼간격이있는n개의숫자(-4, 2))
}

func x만큼간격이있는n개의숫자(x int, n int) []int64 {

	result := make([]int64, n)

	for i := 0; i < n; i++ {
		result[i] = int64(x * (i + 1))
	}

	return result
}

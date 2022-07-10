package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(n2배열자르기(3, 2, 5))
}

func n2배열자르기(n int, left int64, right int64) []float64 {
	return convert2DSliceTo1DSlice(n, left, right)
}

func convert2DSliceTo1DSlice(n int, left int64, right int64) []float64 {

	var result []float64

	for i := left; i <= right; i++ {
		result = append(result, math.Max(float64(int(i)/n), float64(int(i)%n))+1)
	}

	return result
}

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

	for left <= right {
		a := math.Floor(float64(int(left) / n))
		b := float64(int(left+1) % n)
		result = append(result, math.Max(a, b)+1)
		left++
	}

	return result
}

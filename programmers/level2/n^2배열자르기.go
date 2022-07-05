package main

import "fmt"

func main() {
	fmt.Println(n2배열자르기(3, 2, 5))
}

func n2배열자르기(n int, left int64, right int64) []int {
	return getSliceFromLeftToRight(convert2DSliceTo1DSlice(generate2DSlice(n)), left, right)
}

func generate2DSlice(n int) [][]int {

	var result = make([][]int, n, n)

	for i := 0; i < n; i++ {
		result[i] = make([]int, n)

		for j := 0; j < n; j++ {
			result[i][j] = getValueCompareIAndJ(i, j)
		}
	}

	return result
}

func getValueCompareIAndJ(i, j int) int {

	if j > i {
		return j + 1
	}

	return i + 1
}

func convert2DSliceTo1DSlice(sl [][]int) []int {

	var result = make([]int, len(sl))

	for i := 0; i < len(sl); i++ {
		for j := 0; j < len(sl); j++ {
			result[i] = sl[i][j]
		}
	}

	return result
}

func getSliceFromLeftToRight(sl []int, left int64, right int64) []int {

	var result []int

	for i := left; i <= right; i++ {
		result = append(result, sl[i])
	}

	return result
}

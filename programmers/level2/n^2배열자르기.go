package main

func main() {

}

func n2배열자르기(n int, left int64, right int64) []int {
	return getSliceFromLeftToRight(convert2DSliceTo1DSlice(generate2DSlice(n)), left, right)
}

func generate2DSlice(n int) [][]int {

	var result = make([][]int, n, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result[i][j] = i
		}
	}

	return result
}

func convert2DSliceTo1DSlice(sl [][]int) []int {

	var result []int

	for i := 0; i < len(sl); i++ {
		result = append(result, sl[0][i])
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

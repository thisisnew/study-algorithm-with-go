package main

import "fmt"

func main() {
	fmt.Println(n2배열자르기(3, 2, 5))
}

func n2배열자르기(n int, left int64, right int64) []int {
	return convert2DSliceTo1DSlice(n, left, right)
}

func convert2DSliceTo1DSlice(n int, left int64, right int64) []int {

	var sl []int
	var result []int

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			v := getValueCompareIAndJ(i, j)
			sl = append(sl, v)
			if len(sl) >= (int(left)+1) && len(sl) <= (int(right)+1) {
				result = append(result, v)
			}

			if len(result) >= int(right)-int(left)+1 {
				break
			}
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

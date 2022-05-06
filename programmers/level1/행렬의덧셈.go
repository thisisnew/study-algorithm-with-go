package main

import "fmt"

func main() {

	arr1 := [][]int{
		{1, 2},
		{2, 3},
	}

	arr2 := [][]int{
		{3, 4},
		{5, 6},
	}

	fmt.Println(행렬의덧셈(arr1, arr2))
}

func 행렬의덧셈(arr1 [][]int, arr2 [][]int) [][]int {

	result := make([][]int, len(arr1))

	for i := 0; i < len(arr1); i++ {
		result[i] = *sumArray(&arr1[i], &arr2[i])
	}

	return result
}

func sumArray(arr1 *[]int, arr2 *[]int) *[]int {

	result := make([]int, len(*arr1))

	for i := 0; i < len(*arr1); i++ {
		result[i] = (*arr1)[i] + (*arr2)[i]
	}

	return &result
}

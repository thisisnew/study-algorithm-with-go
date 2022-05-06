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

	for i := 0; i < len(arr1); i++ {
		arr1[i] = sumArray(arr1[i], arr2[i])
	}

	return arr1
}

func sumArray(arr1 []int, arr2 []int) []int {

	for i := 0; i < len(arr1); i++ {
		arr1[i] += arr2[i]
	}

	return arr1
}

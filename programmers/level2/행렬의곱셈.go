package main

import "fmt"

func main() {

	arr1 := [][]int{
		{1, 4},
		{3, 2},
		{4, 1},
	}

	arr2 := [][]int{
		{3, 3},
		{3, 3},
	}

	fmt.Println(행렬의곱셈(arr1, arr2))

}

func 행렬의곱셈(arr1 [][]int, arr2 [][]int) [][]int {

	var result [][]int

	for _, ar1 := range arr1 {
		var a []int

		for i, _ := range arr2[0] {

			var sum int

			for j, ar2 := range arr2 {
				sum += ar1[j] * ar2[i]
			}

			a = append(a, sum)

		}

		result = append(result, a)
	}

	return result
}

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

	var idx int
	var result [][]int

	for {

		if idx > len(arr2[0]) {
			break
		}

		for _, ar := range arr1 {

			var a []int

			for _, prop := range ar {
				for _, ar2 := range arr2 {
					a = append(a, prop*ar2[idx])
				}
			}
		}
	}

	return result
}

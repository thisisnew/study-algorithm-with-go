package main

import "fmt"

func main() {
	fmt.Println(genTwoDimension([]int{1, 2, 3, 4, 5, 6, 7, 8}, 2))
}

func genTwoDimension(num_list []int, n int) [][]int {

	var result [][]int
	var dim []int
	var idx = 0
	for i, num := range num_list {

		if i > 0 && i%n == 0 {
			result = append(result, dim)
			dim = []int{}
			idx = 0
		}

		dim = append(dim, num)
		idx++
	}

	if len(dim) > 0 {
		result = append(result, dim)
	}

	return result
}

package main

import (
	"fmt"
	"sort"
)

func main() {

	array := []int{1, 5, 2, 6, 3, 7, 4}
	commands := [][]int{
		{2, 5, 3},
		{4, 4, 1},
		{1, 7, 3},
	}

	fmt.Println(K번째수(array, commands))
}

func K번째수(array []int, commands [][]int) []int {

	var result []int

	for _, cm := range commands {

		cp := make([]int, len(array))
		copy(cp, array)

		partArray := cp[cm[0]-1 : cm[1]]

		sort.Slice(partArray, func(i, j int) bool {
			return partArray[i] < partArray[j]
		})

		result = append(result, partArray[cm[2]-1])

	}

	return result
}

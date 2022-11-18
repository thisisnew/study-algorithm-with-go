package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(진료순서정하기([]int{3, 76, 24}))
}

func 진료순서정하기(emergency []int) []int {

	var temp = make([]int, len(emergency))
	copy(temp, emergency)

	sort.Slice(temp, func(i, j int) bool {
		return temp[i] > temp[j]
	})

	var result = make([]int, len(emergency))
	for i, em := range emergency {
		result[i] = getIndexEmergency(temp, em)
	}

	return result
}

func getIndexEmergency(emergency []int, n int) int {

	for i, em := range emergency {
		if em == n {
			return i + 1
		}
	}

	return -1
}

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(가까운수([]int{1, -1}, 0))
}

func 가까운수(array []int, n int) int {

	var minDif = math.Abs(float64(array[0] - n))
	var min = 0

	for i := 1; i < len(array); i++ {

		dif := math.Abs(float64(array[i] - n))

		if dif < minDif {
			minDif = dif
			min = i
		}

		if dif == minDif {
			if array[i] < array[min] {
				minDif = dif
				min = i
			}
		}
	}

	return array[min]
}

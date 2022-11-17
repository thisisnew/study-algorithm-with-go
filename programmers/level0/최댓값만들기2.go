package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(최댓값만들기2([]int{1, 2, -3, 4, -5}))
}

func 최댓값만들기2(numbers []int) int {

	var abMax = 0

	for _, number := range numbers {
		if int(math.Abs(float64(number))) > abMax {
			abMax = number
		}
	}

	var result = 0
	for _, number := range numbers {
		if abMax == number {
			continue
		}

		if abMax*number > result {
			result = abMax * number
		}
	}

	return result
}

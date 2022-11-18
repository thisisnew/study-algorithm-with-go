package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(분수의덧셈(9, 2, 1, 3))
}

func 분수의덧셈(denum1 int, num1 int, denum2 int, num2 int) []int {

	var total = float64(denum1)/float64(num1) + float64(denum2)/float64(num2)

	st := 1

	for {
		if math.Trunc(total) == total {
			break
		}

		total *= 10
		st *= 10
	}

	result := gcdPoints(total, float64(st))

	return []int{int(total / result), st / int(result)}
}

func gcdPoints(total float64, st float64) float64 {
	if st == 0 {
		return total
	} else {
		return gcdPoints(st, float64(int(total)%int(st)))
	}
}

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(점찍기(1, 5))
}

func 점찍기(k int, d int) int64 {

	var result float64

	for i := 0; i <= d; i += k {
		y := math.Sqrt(math.Pow(float64(d), 2.0) - math.Pow(float64(i), 2.0))
		result += math.Floor(y/float64(k)) + 1
	}

	return int64(result)
}

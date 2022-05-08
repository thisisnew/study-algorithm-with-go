package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(정수제곱근판별(121))
}

func 정수제곱근판별(n int64) int64 {

	var result int64 = -1

	sqrt := math.Sqrt(float64(n))

	if math.Mod(sqrt, 1.0) == 0 {
		result = int64(math.Pow(sqrt+1, 2))
	}

	return result
}

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(정수제곱근판별(121))
}

func 정수제곱근판별(n int64) int64 {

	sqrt := math.Sqrt(float64(n))

	if math.Mod(sqrt, 1.0) == 0 {
		return int64((sqrt + 1) * (sqrt + 1))
	} else {
		return -1
	}

}

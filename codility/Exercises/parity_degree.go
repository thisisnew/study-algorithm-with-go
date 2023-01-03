package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(ParityDegree(24))
}

func ParityDegree(N int) int {

	var result float64 = 0

	for {
		if int(math.Pow(2.0, result)) >= N {
			return int(result)
		}

		result++
	}

}

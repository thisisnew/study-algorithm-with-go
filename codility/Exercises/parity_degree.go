package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(ParityDegree(24))
}

func ParityDegree(N int) int {

	var result = 0
	var n float64 = 0

	for {
		pow := int(math.Pow(2.0, n))

		if pow >= N {
			return result
		}

		if N%(pow) == 0 {
			result = int(n)
		}

		n++
	}

}

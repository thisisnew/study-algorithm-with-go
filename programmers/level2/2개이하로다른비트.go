package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	numbers := []int64{2, 7}

	fmt.Println(두개이하로다른비트(numbers))
}

func 두개이하로다른비트(numbers []int64) []int64 {

	var result = make([]int64, len(numbers))

	for i, n := range numbers {

		if n%2 != 0 {
			result[i] = n + int64(math.Pow(2, countOneBits(n)-1))
		} else {
			result[i] = n + 1
		}

	}

	return result
}

func countOneBits(n int64) float64 {

	var result float64
	var binN = strconv.FormatInt(n, 2)

	for i := len(binN) - 1; i >= 0; i-- {

		if binN[i:i+1] == "0" {
			break
		}

		result++

	}

	return result
}

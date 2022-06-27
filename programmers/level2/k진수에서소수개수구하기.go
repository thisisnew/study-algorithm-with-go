package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(k진수에서소수개수구하기(524287, 2))
}

func k진수에서소수개수구하기(n int, k int) int {

	var result int

	kStringSlice := strings.Split(convertDecimalToKNumber(n, k), "0")

	for _, kString := range kStringSlice {
		result += addResultIfPrimeNumber(kString)
	}

	return result
}

func convertDecimalToKNumber(n int, k int) string {
	return strconv.FormatInt(int64(n), k)
}

func addResultIfPrimeNumber(num string) int {

	n, err := strconv.ParseFloat(num, 64)

	if n <= 1 || err != nil {
		return 0
	}

	sqrt := math.Sqrt(n)

	var i float64 = 2

	for {

		if i > sqrt {
			break
		}

		if int(n)%int(i) == 0 {
			return 0
		}

		i++
	}

	return 1
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(소수찾기("17"))
}

func 소수찾기(numbers string) int {

	var result []int

	numbersRune := []rune(numbers)

	for i := 0; i < len(numbersRune); i++ {
		var num strings.Builder
		num.WriteRune(numbersRune[i])

		for j := 0; j < len(numbersRune); j++ {
			n, _ := strconv.Atoi(num.String())

			if isPrimeNumber(n) {
				result = append(result, n)
			}

			if i == j {
				continue
			}

			num.WriteRune(numbersRune[j])
		}

	}

	fmt.Println(result)

	return len(result)
}

func isPrimeNumber(n int) bool {

	if n == 1 {
		return false
	}

	for i := 2; i < n-1; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

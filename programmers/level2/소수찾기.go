package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(소수찾기("011"))
}

func 소수찾기(numbers string) int {

	var result []int

	numbersRune := []rune(numbers)

	for i := 0; i < len(numbersRune); i++ {
		var num strings.Builder
		num.WriteRune(numbersRune[i])

		ok, n := isPrimeNumber(&num)

		if ok && isDup(n, result) {
			result = append(result, n)
		}

		for j := 0; j < len(numbersRune); j++ {
			if i == j {
				continue
			}

			num.WriteRune(numbersRune[j])

			ok, n := isPrimeNumber(&num)

			if ok && isDup(n, result) {
				result = append(result, n)
			}
		}

	}

	fmt.Println(result)

	return len(result)
}

func isPrimeNumber(num *strings.Builder) (bool, int) {

	n, _ := strconv.Atoi(num.String())

	if n <= 1 {
		return false, -1
	}

	for i := 2; i < n-1; i++ {
		if n%i == 0 {
			return false, -1
		}
	}

	return true, n
}

func isDup(n int, result []int) bool {

	for _, p := range result {
		if p == n {
			return true
		}
	}

	return false
}

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	numbers := []int{3, 30, 34, 5, 9}

	fmt.Println(가장큰수(numbers))
}

func 가장큰수(numbers []int) string {

	digit := 10

	sort.Slice(numbers, func(i, j int) bool {

		ni := getSingleDigit(numbers[i], digit)
		nj := getSingleDigit(numbers[j], digit)

		for {
			if (ni < 10 && nj < 10) || ni != nj {
				break
			}

			digit = digit * 10
			ni = getSingleDigit(numbers[i], digit)
			nj = getSingleDigit(numbers[j], digit)
		}

		return ni > nj
	})

	var result strings.Builder

	for _, n := range numbers {
		result.WriteString(strconv.Itoa(n))
	}

	return result.String()
}

func getSingleDigit(n int, digit int) int {

	for {
		if n/digit < digit {
			return n / digit
		}

		n = n / digit
	}

}

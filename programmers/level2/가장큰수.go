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

		//for {
		//	if ni != nj {
		//		break
		//	}
		//
		//	digit *= 10
		//	ni = getSingleDigit(numbers[i], digit)
		//	nj = getSingleDigit(numbers[j], digit)
		//}

		if ni == nj {
			return numbers[i] > numbers[j]
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

	var result = n

	for {
		if result < digit {
			return result
		}

		result = result / digit
	}

}

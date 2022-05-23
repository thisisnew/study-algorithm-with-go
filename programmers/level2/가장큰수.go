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

	sort.Slice(numbers, func(i, j int) bool {
		a := strconv.Itoa(numbers[i])
		b := strconv.Itoa(numbers[j])

		return a+b > b+a
	})

	if numbers[0] == 0 {
		return "0"
	}

	var result strings.Builder

	for _, n := range numbers {
		result.WriteString(strconv.Itoa(n))
	}

	return result.String()
}

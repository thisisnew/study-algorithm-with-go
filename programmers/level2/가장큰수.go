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

	var result = make([]string, len(numbers))

	for i, n := range numbers {
		result[i] = strconv.Itoa(n)
	}

	sort.Slice(result, func(i, j int) bool {
		a, _ := strconv.Atoi(fmt.Sprintf("%s%s", result[i], result[j]))
		b, _ := strconv.Atoi(fmt.Sprintf("%s%s", result[j], result[i]))

		return a-b >= 0
	})

	if result[0] == "0" {
		return "0"
	}

	return strings.Join(result, "")

}

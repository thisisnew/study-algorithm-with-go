package main

import (
	"fmt"
	"sort"
	"strconv"
	"unicode"
)

func main() {
	fmt.Println(문자열정렬하기1("hi12392"))
}

func 문자열정렬하기1(my_string string) []int {

	var result []int

	for _, s := range my_string {
		if unicode.IsNumber(s) {
			n, _ := strconv.Atoi(string(s))
			result = append(result, n)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return result
}

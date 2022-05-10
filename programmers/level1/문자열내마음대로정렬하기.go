package main

import (
	"fmt"
	"sort"
)

func main() {

	var strings = []string{"sun", "bed", "car"}
	var n = 1

	fmt.Println(문자열내마음대로정렬하기(strings, n))
}

func 문자열내마음대로정렬하기(strings []string, n int) []string {

	sort.Slice(strings, func(i, j int) bool {

		ci := strings[i][n : n+1]
		cj := strings[j][n : n+1]

		if ci == cj {
			return strings[i] < strings[j]
		}

		return ci < cj

	})

	return strings
}

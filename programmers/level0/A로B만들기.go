package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(A로B만들기("olleh", "hello"))
}

func A로B만들기(before string, after string) int {
	b := []rune(before)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	a := []rune(after)
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	if string(b) == string(a) {
		return 1
	}

	return 0
}

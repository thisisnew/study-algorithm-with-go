package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(소수찾기("17"))
}

func 소수찾기(numbers string) int {

	var result int

	numbersRune := []rune(numbers)

	sort.Slice(numbersRune, func(i, j int) bool {
		return numbersRune[i] > numbersRune[j]
	})

	for i := 0; i < len(numbersRune); i++ {
		if numbersRune[i] == '0' {
			continue
		}

	}

	return result
}

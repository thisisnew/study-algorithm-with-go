package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(큰수만들기("1231234", 3))
}

func 큰수만들기(number string, k int) string {

	var ln = len([]rune(number)) - k
	var result = make([]string, ln)
	var idx int

	for i, n := range number {

		if len(strings.TrimSpace(result[idx])) == 0 {
			result[idx] = string(n)
			continue
		}

		if string(n) > result[idx] {
			result[idx] = string(n)
		}

		if len([]rune(number[i:])) <= ln-len(result) {
			idx++
		}
	}

	return strings.Join(result, "")
}

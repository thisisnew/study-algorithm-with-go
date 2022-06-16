package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(큰수만들기("1231234", 3))
}

func 큰수만들기(number string, k int) string {

	var result = []string{number[0:1]}
	var ln = len([]rune(number)) - k
	var idx int

	for i := 1; i < len(number); i++ {
		if number[i:i+1] > result[idx] {
			result[idx] = number[i : i+1]
		}

		if len([]rune(number[i:])) <= ln-len(result) {
			idx++
		}
	}

	return strings.Join(result, "")
}

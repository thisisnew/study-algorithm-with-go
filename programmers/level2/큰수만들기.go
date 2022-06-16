package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(큰수만들기("1924", 3))
}

func 큰수만들기(number string, k int) string {

	var ln = len([]rune(number)) - k
	var result = make([]string, ln)
	var idx int

	for i, num := range number {

		remain := len([]rune(number[i:]))

		if remain <= ln-1-idx {
			idx++
		}

		n := string(num)

		if len(strings.TrimSpace(result[idx])) == 0 {
			result[idx] = n
			continue
		}

		if n > result[idx] {
			result[idx] = n
		}
	}

	return strings.Join(result, "")
}

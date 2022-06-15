package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(큰수만들기("1231234", 3))
}

func 큰수만들기(number string, k int) string {

	var result []int
	var ln = len([]rune(number)) - k
	var idx int

	for i, num := range number {

		n, _ := strconv.Atoi(string(num))

		if len(result) == 0 {
			result = append(result, n)
			continue
		}

		if n > result[idx] {
			result[idx] = n
		}

		if len([]rune(number[i:])) <= ln-len(result) {
			idx++
		}
	}

	return ""
}

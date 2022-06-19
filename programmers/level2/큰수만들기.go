package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(큰수만들기("4177252841", 4))
}

func 큰수만들기(number string, k int) string {

	var result strings.Builder
	var ln = len([]rune(number)) - k
	var start = 0

	for {

		if start >= len([]rune(number)) || len([]rune(number)) == ln {
			break
		}

		left := len([]rune(number)) + k + 1
		max := 0

		for i := start; i < left; i++ {

			num, _ := strconv.Atoi(number[i : i+1])

			if max >= num {
				continue
			}

			max = num
			start = i + 1
		}

		result.WriteString(strconv.Itoa(max))
	}

	return result.String()
}

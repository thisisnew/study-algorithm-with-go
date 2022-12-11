package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(큰수만들기("1924", 2))
}

func 큰수만들기(number string, k int) string {

	var result strings.Builder
	var numbers = strings.Split(number, "")
	var ln = len(numbers) - k

	for {
		if result.Len() == ln {
			return result.String()
		}

		var max = numbers[0]
		var idx = 0

		for i, n := range numbers {
			if n > max {
				max = n
				idx = i
			}
		}

		result.WriteString(max)
		numbers = append(numbers[:idx], numbers[idx+1:]...)
	}
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(큰수만들기("4177252841", 4))
}

func 큰수만들기(number string, k int) string {

	var result strings.Builder
	var numbers = strings.Split(number, "")
	var ln = len(numbers) - k

	for {
		if ln == 0 {
			return result.String()
		}

		var max = numbers[0]
		var idx = 0

		for i := 1; i <= len(numbers)-ln; i++ {
			n := numbers[i]

			if n > max {
				max = n
				idx = i
			}
		}

		ln--
		result.WriteString(max)
		numbers = numbers[idx+1:]
	}
}

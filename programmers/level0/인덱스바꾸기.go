package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(인덱스바꾸기("hello", 1, 2))
}

func 인덱스바꾸기(my_string string, num1 int, num2 int) string {

	var result strings.Builder

	for i, s := range my_string {

		if i == num1 {
			result.WriteString(my_string[num2 : num2+1])
			continue
		}

		if i == num2 {
			result.WriteString(my_string[num1 : num1+1])
			continue
		}

		result.WriteRune(s)
	}

	return result.String()
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(문자반복출력하기("hello", 3))
}

func 문자반복출력하기(my_string string, n int) string {

	var result strings.Builder

	for _, s := range my_string {
		str := string(s)

		for i := 0; i < n; i++ {
			result.WriteString(str)
		}
	}

	return result.String()
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(잘라서배열로저장하기("abc1Addfggg4556b", 6))
}

func 잘라서배열로저장하기(my_str string, n int) []string {

	var result []string
	ln := len([]rune(my_str))
	var sb strings.Builder

	for i := 0; i < ln; i++ {
		if i > 0 && i%n == 0 {
			result = append(result, sb.String())
			sb.Reset()
		}

		sb.WriteString(my_str[i : i+1])
	}

	if sb.Len() > 0 {
		result = append(result, sb.String())
	}

	return result
}

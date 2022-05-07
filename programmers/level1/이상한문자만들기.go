package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(이상한문자만들기("try hello world"))
}

func 이상한문자만들기(s string) string {

	var result []string
	arr := strings.Split(s, " ")

	for _, item := range arr {
		var sb strings.Builder

		for i, ch := range item {
			if (i+1)%2 == 0 {
				sb.WriteString(string(ch))
			} else {
				sb.WriteString(strings.ToUpper(string(ch)))
			}
		}

		result = append(result, sb.String())
	}

	return strings.Join(result, " ")
}

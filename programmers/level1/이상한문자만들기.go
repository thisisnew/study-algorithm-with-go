package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(이상한문자만들기("try hello world"))
}

func 이상한문자만들기(s string) string {

	var idx int
	var sb strings.Builder

	for _, ch := range s {

		if ch == 32 {
			idx = 0
			sb.WriteString(" ")
			continue
		}

		if idx%2 == 0 {
			sb.WriteString(strings.ToUpper(string(ch)))
		} else {
			sb.WriteString(strings.ToLower(string(ch)))
		}

		idx++
	}

	return sb.String()
}

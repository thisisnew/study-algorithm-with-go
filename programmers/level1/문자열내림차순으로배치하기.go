package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(문자열내림차순으로배치하기("Zbcdefg"))
}

func 문자열내림차순으로배치하기(s string) string {

	var result strings.Builder

	for i := len(s) - 1; i >= 0; i-- {
		result.WriteString(s[i : i+1])
	}

	return result.String()
}

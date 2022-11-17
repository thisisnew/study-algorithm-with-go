package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(모음제거("nice to meet you"))
}

func 모음제거(my_string string) string {

	var result strings.Builder

	for _, s := range my_string {

		if unicode.IsSpace(s) {
			result.WriteString(" ")
			continue
		}

		str := string(s)

		switch str {
		case "a", "e", "i", "o", "u":
		default:
			result.WriteString(str)
		}

	}

	return result.String()
}

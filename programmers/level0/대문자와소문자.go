package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(대문자와소문자("cccCCC"))
}

func 대문자와소문자(my_string string) string {

	var result strings.Builder

	for _, s := range my_string {

		str := string(s)

		if unicode.IsUpper(s) {
			result.WriteString(strings.ToLower(str))
		} else {
			result.WriteString(strings.ToUpper(str))
		}

	}

	return result.String()
}

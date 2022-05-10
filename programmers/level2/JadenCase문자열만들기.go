package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(JadenCase문자열만들기("for the last week"))
}

func JadenCase문자열만들기(s string) string {

	var result strings.Builder
	var idx = 0

	for _, st := range s {

		if unicode.IsSpace(st) {
			result.WriteRune(st)
			idx = 0
			continue
		}

		if idx == 0 {
			result.WriteString(strings.ToUpper(string(st)))
		} else {
			result.WriteString(strings.ToLower(string(st)))
		}

		idx++
	}

	return result.String()
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(중복된문자제거("We are the world"))
}

func 중복된문자제거(my_string string) string {

	var result strings.Builder
	var rs []rune

	for _, s := range my_string {

		if isDupRune(s, rs) {
			continue
		}

		result.WriteRune(s)
		rs = append(rs, s)
	}

	return result.String()
}

func isDupRune(r rune, rs []rune) bool {

	for _, p := range rs {
		if p == r {
			return true
		}
	}

	return false
}

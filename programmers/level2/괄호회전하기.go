package main

import "fmt"

func main() {
	fmt.Println(괄호회전하기("[)(]"))
}

func 괄호회전하기(s string) int {

	var result int

	var sCopy = s

	for {

		if isValidBracedString(s) {
			result++
		}

		s = moveTokenLeft(s)

		if s == sCopy {
			break
		}

	}

	return result
}

func moveTokenLeft(s string) string {
	return s[1:] + s[0:1]
}

func isValidBracedString(s string) bool {

	//first
	switch s[0:1] {
	case "}", "]", ")":
		return false
	}

	//last
	switch s[len(s)-1:] {
	case "{", "[", "(":
		return false
	}

	return true
}

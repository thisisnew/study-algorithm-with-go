package main

import "fmt"

func main() {
	fmt.Println(괄호회전하기("[](){}"))
}

func 괄호회전하기(s string) int {

	var result int

	for {

		if isValidBracedString(s) {
			break
		}

	}

	return result
}

func isValidBracedString(s string) bool {

	firstLetter := s[0:1]
	lastLetter := s[len(s)-1:]

	switch firstLetter {
	case "}", "]", ")":
		return false
	}

	switch lastLetter {
	case "{", "[", "(":
		return false
	}

	return true
}

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

	var cnt int

	for _, c := range s {

		switch string(c) {
		case "{", "[", "(":
			cnt++
		case "}", "]", ")":
			cnt--
		}

		if cnt < 0 {
			return false
		}

	}

	return true
}

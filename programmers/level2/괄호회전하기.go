package main

import "fmt"

type sBraceStack []string
type mBraceStack []string
type lBraceStack []string

var sBraceLocked = false
var mBraceLocked = false
var lBraceLocked = false

func main() {
	fmt.Println(괄호회전하기("}}}"))
}

func 괄호회전하기(s string) int {

	var result int
	var origin = s

	for {
		sBraceLocked = false
		mBraceLocked = false
		lBraceLocked = false

		if isValidBraces(s) {
			result++
		}

		s = moveTokenLeft(s)

		if s == origin {
			break
		}
	}

	return result
}

func moveTokenLeft(s string) string {
	return s[1:] + s[0:1]
}

func isValidBraces(bs string) bool {
	for _, s := range bs {
		if isBraceLocked(s) {
			continue
		}

		return false
	}

	return true
}

func isBraceLocked(b rune) bool {

	switch b {
	case '(':
		sBraceLocked = true
	case '{':
		mBraceLocked = true
	case '[':
		lBraceLocked = true
	case ')':
		if !sBraceLocked {
			return false
		}
		sBraceLocked = false
	case '}':
		if !mBraceLocked {
			return false
		}
		mBraceLocked = false
	case ']':
		if !lBraceLocked {
			return false
		}
		lBraceLocked = false
	}

	return true
}

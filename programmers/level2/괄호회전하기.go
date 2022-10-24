package main

import "fmt"

var sBraceStack []string
var mBraceStack []string
var lBraceStack []string

func popBraceStack(bs []string) []string {

	if len(bs) == 0 {
		return []string{}
	}

	return bs[:len(bs)-1]
}

func main() {
	fmt.Println(괄호회전하기("}}}"))
}

func 괄호회전하기(s string) int {

	var result int
	var origin = s

	for {

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
		if isBraceLocked(string(s)) {
			continue
		}

		return false
	}

	return true
}

func isBraceLocked(b string) bool {

	switch b {
	case "(":
		sBraceStack = append(sBraceStack, b)
	case "{":
		mBraceStack = append(mBraceStack, b)
	case "[":
		lBraceStack = append(lBraceStack, b)
	case ")":
		if len(sBraceStack) == 0 {
			return false
		}

		sBraceStack = popBraceStack(sBraceStack)
	case "}":
		if len(mBraceStack) == 0 {
			return false
		}

		mBraceStack = popBraceStack(mBraceStack)
	case "]":
		if len(lBraceStack) == 0 {
			return false
		}

		lBraceStack = popBraceStack(lBraceStack)
	}

	return true
}

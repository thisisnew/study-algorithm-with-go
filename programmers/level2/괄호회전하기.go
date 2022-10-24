package main

import "fmt"

type braceStack []string

var bs braceStack

func (bs *braceStack) pop() string {

	if len(*bs) == 0 {
		return ""
	}

	var result = (*bs)[len(*bs)-1]
	*bs = (*bs)[:len(*bs)-1]

	return result
}

func (bs *braceStack) push(s string) {
	*bs = append(*bs, s)
}

func main() {
	fmt.Println(괄호회전하기("([{)}]"))
}

func 괄호회전하기(s string) int {

	var result int
	var origin = s

	for {
		if isValidBraces(s) && len(bs) == 0 {
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
	case "(", "{", "[":
		bs.push(b)
	case ")":
		p := bs.pop()

		if p != "(" {
			return false
		}
	case "}":
		p := bs.pop()

		if p != "{" {
			return false
		}
	case "]":
		p := bs.pop()

		if p != "[" {
			return false
		}
	}

	return true
}

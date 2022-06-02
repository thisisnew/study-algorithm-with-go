package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	v *list.List
}

func newStack() *Stack {
	return &Stack{list.New()}
}

func (q *Stack) push(v interface{}) {
	q.v.PushBack(v)
}

func (q *Stack) pop() interface{} {
	back := q.v.Back()
	if back == nil {
		return nil
	}

	return q.v.Remove(back)
}

func (q *Stack) peek() interface{} {
	return q.v.Back()
}

func (q *Stack) isEmpty() book {

}

func main() {
	fmt.Println(괄호회전하기("([{)}]"))
}

func 괄호회전하기(s string) int {

	var result int

	var sCopy = s

	for {

		if isValidBracedString(s) {
			result++
		}

		s = moveTokenLeft(s)

		if s != sCopy {
			continue
		}

		break
	}

	return result
}

func moveTokenLeft(s string) string {
	return s[1:] + s[0:1]
}

func isValidBracedString(s string) bool {

	stack := newStack()

	switch s {
	case "[", "{", "(":
		stack.push(s)
	case "]":
		if stack.peek() != "[" {
			return false
		}

		stack.pop()
	case "}":
		if stack.peek() != "{" {
			return false
		}

		stack.pop()
	case ")":
		if stack.peek() != "(" {
			return false
		}

		stack.pop()
	}

	return true
}

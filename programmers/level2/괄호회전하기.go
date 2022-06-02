package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	v *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (q *Stack) Push(v interface{}) {
	q.v.PushBack(v)
}

func (q *Stack) Pop() interface{} {
	back := q.v.Back()
	if back == nil {
		return nil
	}

	return q.v.Remove(back)
}

func (q *Stack) Peek() interface{} {
	return q.v.Back()
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

	stack := NewStack()

	switch s {
	case "[", "{", "(":
		stack.Push(s)
	case "]":
		if stack.Peek() != "[" {
			return false
		}

		stack.Pop()
	case "}":
		if stack.Peek() != "{" {
			return false
		}

		stack.Pop()
	case ")":
		if stack.Peek() != "(" {
			return false
		}

		stack.Pop()
	}

	return true
}

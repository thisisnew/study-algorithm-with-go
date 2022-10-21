package main

import "fmt"

func main() {
	fmt.Println(괄호회전하기("([{)}]"))
}

func 괄호회전하기(s string) int {

	var result int
	var origin = s

	for {
		s = moveTokenLeft(s)

		if s == origin {
			break
		}

		//검증로직
	}

	return result
}

func moveTokenLeft(s string) string {
	return s[1:] + s[0:1]
}

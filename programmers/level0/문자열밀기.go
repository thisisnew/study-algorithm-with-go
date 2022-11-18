package main

import "fmt"

func main() {
	fmt.Println(문자열밀기("hello", "ohell"))
}

func 문자열밀기(A string, B string) int {

	var limit = len([]rune(A)) + len([]rune(B))
	var cnt = 0

	for {
		if A == B {
			return cnt
		}

		if cnt == limit {
			return -1
		}

		A = fmt.Sprintf("%s%s", A[len([]rune(A))-1:len([]rune(A))], A[0:len([]rune(A))-1])
		cnt++
	}

}

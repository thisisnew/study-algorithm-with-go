package main

import "fmt"

func main() {
	fmt.Println(solution("abcde"))
}

func solution(s string) string {

	var result string
	ln := len(s)

	if ln%2 == 0 {
		result = s[ln/2-1 : ln/2+1]
	} else {
		result = s[ln/2 : ln/2+1]
	}

	return result
}

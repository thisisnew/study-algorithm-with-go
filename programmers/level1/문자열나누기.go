package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(문자열나누기("banana"))
}

func 문자열나누기(s string) int {

	var result []string
	var x string
	var xCnt int
	var yCnt int
	var temp strings.Builder

	for _, r := range s {
		temp.WriteRune(r)

		if x == "" {
			x = string(r)
			continue
		}

		if string(r) == x {
			xCnt++
		} else {
			yCnt++
		}

	}

	return len(result)
}

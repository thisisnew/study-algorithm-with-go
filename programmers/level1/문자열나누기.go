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

		c := string(r)

		if x == "" {
			x = c
		}

		if c == x {
			xCnt++
		} else {
			yCnt++
		}

		if xCnt == yCnt {
			result = append(result, temp.String())
			xCnt = 0
			yCnt = 0
			x = ""
		}

	}

	return len(result)
}

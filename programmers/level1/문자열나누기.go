package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(문자열나누기("aaabbaccccabba"))
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
		}

		switch x {
		case string(r):
			xCnt++
		default:
			yCnt++
		}

		if xCnt != yCnt {
			continue
		}

		result = append(result, temp.String())
		temp.Reset()
		xCnt = 0
		yCnt = 0
		x = ""
	}

	if temp.Len() > 0 {
		result = append(result, temp.String())
	}

	return len(result)
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(StrSymmetryPoint("racecar"))
}

func StrSymmetryPoint(S string) int {

	var ln = len([]rune(S))

	switch ln {
	case 0:
		return -1
	case 1:
		return 0
	}

	var mid = len([]rune(S)) / 2

	if S[0:mid+1] != getReverserStr(mid, S) {
		return -1
	}

	return mid
}

func getReverserStr(idx int, S string) string {

	var result strings.Builder

	for i := len([]rune(S)) - 1; i >= idx; i-- {
		result.WriteString(S[i : i+1])
	}

	return result.String()
}

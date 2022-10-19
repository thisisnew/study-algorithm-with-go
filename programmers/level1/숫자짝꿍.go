package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(숫자짝꿍("100", "2345"))
}

func 숫자짝꿍(X string, Y string) string {

	if len([]rune(X)) > len([]rune(Y)) {
		X, Y = Y, X
	}

	var sl []string
	var yr = []rune(Y)
	var isAllZero = true

	for _, s := range X {
		ok, nYr := isContainsNumberInY(s, yr)

		if ok {
			prop := string(s)

			if isAllZero && prop != "0" {
				isAllZero = false
			}

			sl = append(sl, prop)
			yr = nYr
		}
	}

	if len(sl) == 0 {
		return "-1"
	}

	if isAllZero {
		return "0"
	}

	sort.Slice(sl, func(i, j int) bool {
		return sl[i] > sl[j]
	})

	return strings.Join(sl, "")
}

func isContainsNumberInY(s rune, yr []rune) (bool, []rune) {

	for i, r := range yr {
		if r == s {
			return true, append(yr[:i], yr[i+1:]...)
		}
	}

	return false, yr
}

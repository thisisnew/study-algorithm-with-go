package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(문자열안에문자열("ab6CDE443fgh22iJKlmn1o", "6CD"))
}

func 문자열안에문자열(str1 string, str2 string) int {

	ln := len([]rune(str2))

	for i, s := range str1 {

		var test strings.Builder
		str := string(s)

		if str == str2[0:1] {
			for j := i; j < len([]rune(str1)); j++ {
				test.WriteString(str1[j : j+1])

				if test.Len() == ln {
					if test.String() == str2 {
						return 1
					}

					break
				}
			}
		}
	}

	return 2
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(특정문자제거하기("abcdef", "f"))
}

func 특정문자제거하기(my_string string, letter string) string {

	var result strings.Builder

	for _, s := range my_string {

		str := string(s)

		if str == letter {
			continue
		}

		result.WriteString(str)
	}

	return result.String()
}

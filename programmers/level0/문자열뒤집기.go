package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(문자열뒤집기("jaron"))
}

func 문자열뒤집기(my_string string) string {

	var result strings.Builder

	for i := len([]rune(my_string)) - 1; i >= 0; i-- {
		result.WriteString(my_string[i : i+1])
	}

	return result.String()
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(암호해독("dfjardstddetckdaccccdegk", 4))
}

func 암호해독(cipher string, code int) string {

	var result strings.Builder

	for i := code - 1; i < len([]rune(cipher)); i += code {
		c := cipher[i : i+1]
		result.WriteString(c)
	}

	return result.String()
}

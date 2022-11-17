package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(외계행성의나이(23))
}

func 외계행성의나이(age int) string {

	var alphabets = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "u", "v", "w", "x", "y", "z"}
	var result strings.Builder

	ageStr := strconv.Itoa(age)
	for _, s := range ageStr {
		idx, _ := strconv.Atoi(string(s))
		result.WriteString(alphabets[idx])
	}

	return result.String()
}

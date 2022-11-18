package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(한번만등장한문자("abcabcadc"))
}

func 한번만등장한문자(s string) string {

	var sb strings.Builder

out:
	for i := 0; i < len([]rune(s)); i++ {

		c := s[i : i+1]

		for j := 0; j < len([]rune(s)); j++ {

			if i != j && c == s[j:j+1] {
				continue out
			}
		}

		sb.WriteString(c)
	}

	result := []rune(sb.String())
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return string(result)
}

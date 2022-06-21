package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(큰수만들기("1924", 2))
}

func 큰수만들기(number string, k int) string {

	var result []string

	for _, n := range number {
		for {
			if len(result) == 0 || k <= 0 || result[len(result)-1] >= string(n) {
				break
			}

			result = result[:len(result)-1]
			k--
		}
		result = append(result, string(n))
	}

	return strings.Join(result, "")[0 : len([]rune(number))-k]
}

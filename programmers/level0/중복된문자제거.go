package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(중복된문자제거("people"))
}

func 중복된문자제거(my_string string) string {

	var result strings.Builder
	sl := []rune(my_string)
	var idx = 0

	for {
		if idx >= len(sl) {
			break
		}

		r := sl[idx]
		result.WriteRune(r)

		for i := idx + 1; i < len(sl); i++ {
			if r == sl[i] {
				var temp []rune
				temp = append(sl[0:i], sl[i+1:]...)
				sl = temp
			}
		}

		idx++
	}

	return result.String()
}

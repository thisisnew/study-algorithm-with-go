package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(옹알이1([]string{"ayaye", "uuuma", "ye", "yemawoo", "ayaa"}))
}

func 옹알이1(babbling []string) int {

	var result int

	for _, b := range babbling {
		var sb strings.Builder
		var prev string

		for _, s := range b {
			sb.WriteRune(s)
			st := sb.String()

			if prev != st && canBabbling(st) {
				prev = st
				sb.Reset()
			}
		}

		if sb.Len() == 0 {
			result++
		}
	}

	return result
}

func canBabbling(word string) bool {

	for _, b := range []string{"aya", "ye", "woo", "ma"} {
		if b == word {
			return true
		}
	}

	return false
}

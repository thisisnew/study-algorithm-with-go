package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(옹알이2([]string{"ayaye", "uuu", "yeye", "yemawoo", "ayaayaa"}))
}

func 옹알이2(babbling []string) int {

	var result int

	for _, b := range babbling {
		if canBabbling(b) {
			result++
		}
	}

	return result
}

func canBabbling(bs string) bool {
	var possible = []string{"aya", "ye", "woo", "ma"}

	for _, p := range possible {
		var temp strings.Builder
		var isSameBeginning = false

		for _, b := range bs {
			if !isSameBeginning && string(b) != p[0:1] {
				continue
			}

			isSameBeginning = true
			temp.WriteRune(b)

			if len([]rune(temp.String())) == len([]rune(p)) {

			}

		}
	}

	return false
}

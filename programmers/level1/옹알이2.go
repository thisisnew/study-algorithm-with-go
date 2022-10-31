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

	for _, bab := range babbling {
		var temp strings.Builder
		var prev string

		for _, b := range bab {
			temp.WriteRune(b)

			ts := temp.String()
			if prev != ts && canBabbling(ts) {
				prev = temp.String()
				temp.Reset()
			}
		}

		if temp.Len() == 0 {
			result++
		}
	}

	return result
}

func canBabbling(ts string) bool {
	var possible = []string{"aya", "ye", "woo", "ma"}

	for _, p := range possible {
		if p == ts {
			return true
		}
	}

	return false
}

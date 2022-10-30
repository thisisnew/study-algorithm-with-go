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

		for i := 0; i < len([]rune(bs))-1; i++ {

			if bs[i:i+1] != p[0:1] {
				continue
			}

			temp.WriteString(bs[i : i+1])

		in:
			for j := i + 1; j < len([]rune(bs)); j++ {
				temp.WriteString(bs[j : j+1])

				if len([]rune(temp.String())) == len([]rune(p)) {
					if temp.String() == p {
						temp.Reset()
						continue in
					}

					temp.Reset()
				}
			}
		}
	}

	return false
}

package main

import (
	"fmt"
	"strings"
	"unicode"
)

const (
	CapitalLetterZ = 90
	SmallLetterZ   = 122
)

func main() {
	fmt.Println(시저암호("a B z", 4))
}

func 시저암호(s string, n int) string {

	var result strings.Builder

	for _, rn := range s {

		if unicode.IsSpace(rn) {
			result.WriteString(" ")
			continue
		}

		var ch = rn + int32(n)

		switch {
		case unicode.IsLower(rn):
			if ch > SmallLetterZ {
				ch = ch - int32(26)
			}

		default:
			if ch > CapitalLetterZ {
				ch = ch - int32(26)
			}
		}

		result.WriteRune(ch)
	}

	return result.String()
}

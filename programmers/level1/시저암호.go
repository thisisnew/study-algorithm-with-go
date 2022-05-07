package main

import (
	"fmt"
	"strings"
	"unicode"
)

const (
	CapitalLetterZ int32 = 90
	SmallLetterZ   int32 = 122
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

		var ch int32

		switch {
		case unicode.IsLower(rn):
			ch = moveLetter(rn+int32(n), SmallLetterZ)
		default:
			ch = moveLetter(rn+int32(n), CapitalLetterZ)
		}

		result.WriteRune(ch)
	}

	return result.String()
}

func moveLetter(ch, last int32) int32 {

	if ch <= last {
		return ch
	}

	return ch - int32(26)
}

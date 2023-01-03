package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ThreeLetters(1, 4))
}

func ThreeLetters(A int, B int) string {

	var a = "a"
	var b = "b"
	var total = A + B

	if B > A {
		A, B = B, A
		a, b = b, a
	}

	var result strings.Builder
	var aLp = 0
	var bLp = 0

	for i := 0; i < total; i++ {
		switch {
		case aLp < 2 && A > 0:
			A--
			aLp++
			bLp = 0
			result.WriteString(a)
		case bLp < 2 && B > 0:
			B--
			bLp++
			aLp = 0
			result.WriteString(b)
		}
	}

	return result.String()
}

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
	for i := 0; i < total; i++ {
		if aLp < 2 {
			result.WriteString(a)
			aLp++
			continue
		}

		aLp = 0
		var bLp = 0
		for B > 0 && bLp < 2 {
			bLp++
			B--
			result.WriteString(b)
		}
	}

	return result.String()
}

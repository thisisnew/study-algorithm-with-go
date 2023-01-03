package main

import (
	"fmt"
)

func main() {
	fmt.Println(ThreeLetters(4, 4))
}

func ThreeLetters(A int, B int) string {

	var a = "a"
	var b = "b"
	var total = A + B

	if B > A {
		A, B = B, A
		a, b = b, a
	}

	var result string
	var aLp = 0
	var bLp = 2

	for i := 0; i < total; i++ {
		switch {
		case aLp < 2 && bLp >= 2 && A > 0:
			A--
			aLp++
			result += a
			if aLp == 2 {
				bLp = 0
			}
		case bLp < 2 && aLp >= 2 && B > 0:
			B--
			bLp++
			aLp = 0
			result += b
			if aLp == 2 {
				bLp = 0
			}
		}
	}

	return result
}

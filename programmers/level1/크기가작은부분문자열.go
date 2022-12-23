package main

import "fmt"

func main() {
	fmt.Println(크기가작은부분문자열("3141592", "271"))
}

func 크기가작은부분문자열(t string, p string) int {

	var result int
	var ln = len([]rune(t))
	var cLen = len([]rune(p))

	for i := 0; i < ln; i++ {

		if i > ln-cLen {
			break
		}

		s := t[i : i+cLen]

		if s < p {
			result++
		}
	}

	return result
}

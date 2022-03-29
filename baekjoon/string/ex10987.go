package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var input string
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscan(read, &input)

	var result int

	for _, ch := range input {
		if isVerb(ch) {
			result++
		}
	}

	fmt.Println(result)
}

func isVerb(ch rune) bool {

	var verbs = []rune{'a', 'e', 'i', 'o', 'u'}

	for _, verb := range verbs {
		if verb == ch {
			return true
		}
	}

	return false
}

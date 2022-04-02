package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var input string
	var read = bufio.NewReader(os.Stdin)

	fmt.Fscanln(read, &input)

	var result string

	for _, s := range []rune(input) {

		if isContainsInCAMBRIDGE(s) {
			continue
		}

		result += string(s)
	}

	fmt.Println(result)
}

func isContainsInCAMBRIDGE(r rune) bool {

	for _, rn := range []rune("CAMBRIDGE") {
		if rn == r {
			return true
		}
	}

	return false
}

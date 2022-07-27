package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var read = bufio.NewReader(os.Stdin)

	for {
		text, _, _ := read.ReadLine()

		var input = string(text)

		if input == "END" {
			break
		}

		if hasDuplicateWord(input) {
			continue
		}

		fmt.Println(input)
	}

}

func hasDuplicateWord(input string) bool {

	var sl = make([]bool, int('Z'))

	for _, w := range strings.TrimSpace(input) {
		if sl[w] {
			return true
		}

		sl[w] = true
	}

	return false
}

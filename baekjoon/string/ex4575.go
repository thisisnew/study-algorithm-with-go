package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
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

	var sl = make([]bool, int('Z')+1)

	for _, w := range input {

		if unicode.IsSpace(w) {
			continue
		}

		if sl[w] {
			return true
		}

		sl[w] = true
	}

	return false
}

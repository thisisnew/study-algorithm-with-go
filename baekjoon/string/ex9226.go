package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var read = bufio.NewReader(os.Stdin)

out:
	for {
		var input string
		fmt.Fscanln(read, &input)

		if input == "#" {
			break
		}

		var result = input
		for _, ch := range input {

			if isVowel(ch) {
				fmt.Println(result)

				continue out
			} else {

			}

		}

	}
}

func isVowel(ch rune) bool {

	str := string(ch)

	if str == "a" || str == "e" || str == "i" || str == "o" || str == "u" {
		return true
	}

	return false

}

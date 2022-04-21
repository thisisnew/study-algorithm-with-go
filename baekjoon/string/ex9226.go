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

			isVowel, str := isVowel(ch)

			if isVowel {
				fmt.Println(result + "ay")

				continue out
			} else {
				result = result[1:] + str
			}

		}

	}
}

func isVowel(ch rune) (bool, string) {

	str := string(ch)

	if str == "a" || str == "e" || str == "i" || str == "o" || str == "u" {
		return true, str
	}

	return false, str

}

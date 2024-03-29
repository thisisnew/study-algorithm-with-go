package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &n)

	var word string
	fmt.Fscanln(read, &word)

	var words = []string{word}
	var std = word

	for i := 1; i < n; i++ {
		var word string
		fmt.Fscanln(read, &word)

		if !isSameWord(std, word) {

			if !isContainsInWords(words, word) {
				words = append(words, word)
			}

			std = word
		}
	}

	fmt.Println(len(words))
	return
}

func isSameWord(std, word string) bool {

	if len([]rune(std)) != len([]rune(word)) {
		return false
	}

	var temp = word
	for i := 0; i < len([]rune(word)); i++ {

		if word[i:i+1] != std[0:1] {
			continue
		}

		temp = temp[i:] + temp[0:i]

		if std == temp {
			return true
		}
	}

	return false
}

func isContainsInWords(words []string, word string) bool {

	for _, w := range words {
		if w == word {
			return true
		}
	}

	return false
}

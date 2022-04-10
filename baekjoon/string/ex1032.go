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

	var inputRunes []rune

	for i := 0; i < n; i++ {
		var input string
		fmt.Fscanln(read, &input)

		inputRunes = compareRunes(inputRunes, input)
	}

	fmt.Println(string(inputRunes))
}

func compareRunes(inputRunes []rune, input string) []rune {

	if len(inputRunes) == 0 {
		return []rune(input)
	}

}

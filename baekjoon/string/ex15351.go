package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &n)

	for i := 0; i < n; i++ {
		text, _, _ := read.ReadLine()

		input := strings.Trim(string(text), " ")

		isPerfectLife, result := calculateWord(input)

		if isPerfectLife {
			fmt.Println("PERFECT LIFE")
		} else {
			fmt.Println(result)
		}
	}
}

func calculateWord(input string) (bool, int) {

	var isPerfectLife bool
	var result int

	for _, ch := range input {

	}

	return isPerfectLife, result
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var input string
	var read = bufio.NewReader(os.Stdin)

	for {
		text, _, _ := read.ReadLine()

		input = string(text)

		if input == "END" {
			break
		}

		for i := len(input) - 1; i >= 0; i-- {
			fmt.Print(input[i : i+1])
		}

		fmt.Println()
	}
}

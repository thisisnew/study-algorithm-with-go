package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var read = bufio.NewReader(os.Stdin)

	for {
		var input string
		fmt.Fscanln(read, &input)

		if input == "END" {
			break
		}

		prevRn := input[0:1]

		for i := 1; i < len(input); i++ {
			ch := input[i : i+1]

			if ch == "a" || ch == "e" || ch == "i" || ch == "o" || ch == "u" {

			} else {

			}

		}

	}

}

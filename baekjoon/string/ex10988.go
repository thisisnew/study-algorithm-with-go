package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input string
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscan(read, &input)

	var reverseInput string

	for i := len(input) - 1; i >= 0; i-- {
		reverseInput += input[i : i+1]
	}

	if reverseInput == input {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}

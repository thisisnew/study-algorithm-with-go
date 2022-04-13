package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var read = bufio.NewReader(os.Stdin)
	text, _, _ := read.ReadLine()

	var ucpc = []rune{'U', 'C', 'P', 'C'}
	var index = 0

	for _, t := range string(text) {
		if t == ucpc[index] {
			index++
		}

		if index == 4 {
			fmt.Println("I love UCPC")
			return
		}
	}

	fmt.Println("I hate UCPC")
}

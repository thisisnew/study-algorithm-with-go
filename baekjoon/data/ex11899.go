package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var read = bufio.NewReader(os.Stdin)
	var s string
	fmt.Fscanln(read, &s)

	var right int
	var result int

	for _, b := range s {
		if b == ')' {
			right--
		}

	}

	fmt.Println(result)
}

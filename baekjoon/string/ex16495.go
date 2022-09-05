package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var read = bufio.NewReader(os.Stdin)
	var input string

	fmt.Fscanln(read, &input)

	rns := []rune(input)

	var result float64

	for i := 0; i < len(rns); i++ {
		result *= 26
		result += float64(rns[i]-'A') + 1
	}

	fmt.Println(result)
}

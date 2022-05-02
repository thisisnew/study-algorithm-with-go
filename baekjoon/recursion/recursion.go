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

	fmt.Println(factorial(n))
}

func factorial(n int) int {

	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

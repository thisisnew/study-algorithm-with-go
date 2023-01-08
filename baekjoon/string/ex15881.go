package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	var input string
	var read = bufio.NewReader(os.Stdin)

	fmt.Fscanln(read, &n)
	fmt.Fscanln(read, &input)

}

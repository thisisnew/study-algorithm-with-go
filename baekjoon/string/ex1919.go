package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var a string
	var b string
	var read = bufio.NewReader(os.Stdin)

	fmt.Fscanln(read, &a)
	fmt.Fscanln(read, &b)

}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var read = bufio.NewReader(os.Stdin)
	var n, m int
	var f string
	var s string
	fmt.Fscanln(read, &n, &m)
	fmt.Fscanln(read, &s)
	fmt.Fscanln(read, &f)

	var result int

	fmt.Println(result)
}

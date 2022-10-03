package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var read = bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscanln(read, &n, &m)

	loops := n + m
	for i := 0; i < loops; i++ {
		var url string
		var pw string
		fmt.Fscanln(read, &url, &pw)

	}

}

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

	var loops = 2*(n-1) + 1
	var stars = make([][]string, loops)

	for i := 0; i < loops; i++ {
		stars[i] = make([]string, loops)
		for j := 0; j < loops; j++ {
			stars[i][j] = " "
		}
	}

}

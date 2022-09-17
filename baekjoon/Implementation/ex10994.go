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

	var max = 4*(n-1) + 1

	if max == 1 {
		fmt.Println("*")
		return
	}

	for i := 0; i < max; i++ {

	}

}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var n int
	var read = bufio.NewReader(os.Args)
	fmt.Fscanln(read, &n)

	for i := 0; i < n; i++ {
		var num int
		fmt.Fscanln(read, &num)

	}

}

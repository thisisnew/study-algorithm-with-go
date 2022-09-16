package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &n)

	for i := 0; i < n; i++ {

		var sl = make([]string, n)

		for j := 0; j < n; j++ {
			sl[j] = "*"
		}
		if i%2 > 0 {
			fmt.Print(" ")
		}

		fmt.Print(strings.Join(sl, " "))

		if i < n-1 {
			fmt.Println()
		}
	}
}

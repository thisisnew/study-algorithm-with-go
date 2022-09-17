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

		var result strings.Builder

		if i%2 > 0 {
			result.WriteString(" ")
		}

		for j := 0; j < n; j++ {
			result.WriteString("*")

			if j < n-1 {
				result.WriteString(" ")
			}
		}

		fmt.Print(result.String())

		if i < n-1 {
			fmt.Println()
		}
	}
}

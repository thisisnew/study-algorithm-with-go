package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var read = bufio.NewReader(os.Stdin)
	var s string
	fmt.Fscanln(read, &s)

	var cnt int
	var result int

	for _, b := range s {

		if b == ')' {
			if cnt <= 0 {
				result++
			}
			cnt--
		}

		if b == '(' {
			cnt++
		}

	}

	fmt.Println(result + cnt)
}

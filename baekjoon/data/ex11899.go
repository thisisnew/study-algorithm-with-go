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
	var isClose = false

	for _, b := range s {

		if b == ')' {
			if !isClose {
				cnt--
			}

			cnt--
			if cnt <= 0 {
				result++
			}
		}

		if b == '(' {
			isClose = true
			cnt++
		}

	}

	fmt.Println(result + cnt)
}

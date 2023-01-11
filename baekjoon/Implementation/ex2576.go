package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var read = bufio.NewReader(os.Stdin)
	var sum int
	var min int

	for {
		var n int
		fmt.Fscanln(read, &n)

		if n == 0 {
			break
		}

		if n%2 > 0 {
			sum += n
			if min == 0 || n < min {
				min = n
			}
		}

	}

	fmt.Println(sum)
	fmt.Println(min)
}

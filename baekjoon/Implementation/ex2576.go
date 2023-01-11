package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var read = bufio.NewReader(os.Stdin)
	var sum int
	var min int

	for {
		var n int
		_, err := fmt.Fscanln(read, &n)

		if err == io.EOF {
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

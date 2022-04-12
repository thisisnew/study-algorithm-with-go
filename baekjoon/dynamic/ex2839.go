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

	var result int

	for {
		if n%5 != 0 {
			if n <= 0 {
				break
			}

			n -= 3
			result++
		}
	}

	if n > 0 {
		result += n / 5
	} else if n < 0 {
		result = -1
	}

	fmt.Println(result)
}

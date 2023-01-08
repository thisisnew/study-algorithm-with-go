package main

import (
	"bufio"
	"fmt"
	"os"
)

const penPineAppleApplePen = "pPAp"

func main() {
	var n int
	var input string
	var read = bufio.NewReader(os.Stdin)

	fmt.Fscanln(read, &n)
	fmt.Fscanln(read, &input)

	var i = 0
	var result = 0

	for {
		if i >= n-3 || n < 4 {
			break
		}

		if input[i:i+4] != penPineAppleApplePen {
			i++
			continue
		}

		result++
		i += 3
	}

	fmt.Println(result)
	return
}

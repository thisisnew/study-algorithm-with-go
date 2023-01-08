package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const penPineAppleApplePen = "pPAp"

func main() {
	var n int
	var input string
	var read = bufio.NewReader(os.Stdin)

	fmt.Fscanln(read, &n)
	fmt.Fscanln(read, &input)

	var result = 0
	var f = -1

	for {
		f = strings.Index(input, penPineAppleApplePen)

		if f == -1 {
			break
		}

		result++
		input = input[f+4:]
	}

	fmt.Println(result)
	return
}

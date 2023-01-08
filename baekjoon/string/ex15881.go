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

	var i = 0
	var result = 0
	var sb strings.Builder

	for {

		if i >= n-3 || n < 4 {
			break
		}

		if input[i:i+4] != penPineAppleApplePen {
			i++
			continue
		}

		input = input[0:i] + input[i+4:]
		i = 0
		sb.Reset()
		result++
		n = len([]rune(input))
	}

	fmt.Println(result)
}

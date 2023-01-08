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

		if i > len([]rune(input))-3 || len([]rune(input)) < 4 {
			break
		}

		s := input[i : i+1]

		if s != "p" {
			i++
			continue
		}

		for j := i; j < i+4; j++ {
			sb.WriteString(input[j : j+1])
		}

		if sb.String() != penPineAppleApplePen {
			i++
			continue
		}

		input = input[0:i] + input[i+4:]
		i = 0
		sb.Reset()
		result++
	}

	fmt.Println(result)
}

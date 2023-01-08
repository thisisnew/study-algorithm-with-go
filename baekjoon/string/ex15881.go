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

	var idx = 0
	var f = -1
	var result = 0
	var sb strings.Builder

	for {
		s := input[idx : idx+1]

		if s != "p" {
			continue
		}
		f = idx
		sb.WriteString(input[idx : idx+1])

		if idx < 0 || len([]rune(input)) < 4 {
			fmt.Println(result)
			return
		}

		result++
		input = input[0:f] + input[f+4:]
		idx = 0
	}

}

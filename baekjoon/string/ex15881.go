package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	var input string
	var read = bufio.NewReader(os.Stdin)

	fmt.Fscanln(read, &n)
	fmt.Fscanln(read, &input)

	var idx = -1
	var result = 0

	for {
		idx = strings.Index(input, "pPAp")

		if idx < 0 || len([]rune(input)) < 4 {
			fmt.Println(result)
			return
		}

		result++
		input = input[0:idx] + input[idx+4:]
	}

}

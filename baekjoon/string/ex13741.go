package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var input string
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &input)

	var index = int('a')
	var letters strings.Builder

	for _, c := range input {

		if index != int(c) {
			continue
		}

		letters.WriteString(string(c))
		index++
	}

	fmt.Println(26 - len([]rune(letters.String())))
}

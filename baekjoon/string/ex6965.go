package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var read = bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(read, &n)

	for i := 0; i < n; i++ {
		input, _, _ := read.ReadLine()
		fmt.Println(replace4lengthWordToAsterisk(string(input)))

		if i != n-1 {
			fmt.Println()
		}
	}
}

func replace4lengthWordToAsterisk(input string) string {

	tokens := strings.Split(input, " ")
	var result = make([]string, len(tokens))

	for i, token := range tokens {
		if len([]rune(token)) == 4 {
			result[i] = "****"
			continue
		}

		result[i] = token
	}

	return strings.Join(result, " ")
}

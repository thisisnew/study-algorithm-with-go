package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var read = bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscanln(read, &t)

	for i := 0; i < t; i++ {
		var input string
		fmt.Fscanln(read, &input)

		f := getRepeatString(getTsyStack(input))

		fmt.Println(f)
	}

}

func getTsyStack(input string) []rune {
	return []rune(input)
}

func getRepeatString(tsyStack []rune) string {

	if len(tsyStack) <= 1 {
		return string(tsyStack[0])
	}

	var fResult strings.Builder
	f := tsyStack[0]
	var i = 0

	for {

		if i%2 != 0 {
			i++
			continue
		}

		r := tsyStack[0]

		if r == f {
			return fResult.String()
		}

		fResult.WriteRune(tsyStack[0])
		tsyStack = append(tsyStack[1:], f)
		i++
	}

	return fResult.String()
}

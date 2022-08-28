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

		f, s := getRepeatString(getTsyStack(input))

		fmt.Println(f)
		fmt.Println(s)
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

	for {

		r := tsyStack[0]

		if r == f {
			return fResult.String()
		}

		fResult.WriteRune(tsyStack[0])

		tsyStack = append(...tsyStack[1:], tsyStack[0])
	}

}

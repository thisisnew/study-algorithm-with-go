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

		f := getRepeatString(getTsyStack(input), true)
		s := getRepeatString(getTsyStack(input), false)
		fmt.Println(f)
		fmt.Println(s)
	}

}

func getTsyStack(input string) []rune {
	return []rune(input)
}

func getRepeatString(tsyStack []rune, isFirst bool) string {

	if len(tsyStack) <= 1 {
		return string(tsyStack[0])
	}

	var fResult strings.Builder

	if isFirst {

		for i := 0; i < len(tsyStack); i += 2 {
			fResult.WriteRune(tsyStack[i])
		}

		for i := 1; i < len(tsyStack); i += 2 {
			fResult.WriteRune(tsyStack[i])
		}

	} else {
		for i := 0; i < len(tsyStack); i += 2 {
			fResult.WriteRune(tsyStack[i])
		}

		for i := 1; i < len(tsyStack); i += 2 {
			fResult.WriteRune(tsyStack[i])
		}
	}

	return fResult.String()
}

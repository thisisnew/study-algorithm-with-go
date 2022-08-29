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

	var result string

	switch {
	case len(tsyStack)%2 == 0:

	case len(tsyStack)%2 != 0:
		result = getRepeatStringByOddNumber(tsyStack, isFirst)
	}

	return result
}

func getRepeatStringByOddNumber(tsyStack []rune, isFirst bool) string {

	var result strings.Builder

	switch {
	case isFirst:
		result.WriteString(getStringEvenIndex(tsyStack))
		result.WriteString(getStringOddIndex(tsyStack))

	default:
		result.WriteString(getStringOddIndex(tsyStack))
		result.WriteString(getStringEvenIndex(tsyStack))
	}

	return result.String()
}

func getStringEvenIndex(tsyStack []rune) string {

	var result strings.Builder

	for i := 0; i < len(tsyStack); i += 2 {
		result.WriteRune(tsyStack[i])
	}

	return result.String()
}

func getStringOddIndex(tsyStack []rune) string {

	var result strings.Builder

	for i := 1; i < len(tsyStack); i += 2 {
		result.WriteRune(tsyStack[i])
	}

	return result.String()
}

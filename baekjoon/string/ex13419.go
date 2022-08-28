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

func getRepeatString(tsyStack []rune) (string, string) {

	var f strings.Builder
	var s strings.Builder

}

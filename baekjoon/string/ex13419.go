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
		fmt.Println(getFirstRepeatString(input))
		fmt.Println(getSecondRepeatString(input))
	}

}

func getFirstRepeatString(input string) string {

	if len([]rune(input)) <= 1 {
		return input
	}

	var result strings.Builder
	var first = input[0:1]
	var s = input[0:1]
	var i = 0
	var start = false

	for {
		if i%2 != 0 {
			i++
			continue
		}

		s = input[i : i+1]

		if start && s == first {
			break
		}

		i++

		if i > len([]rune(input))-1 {
			i = 0
		}

		result.WriteString(s)

		if !start {
			start = true
		}
	}

	return result.String()
}

func getSecondRepeatString(input string) string {

	if len([]rune(input)) <= 1 {
		return input
	}

	var result strings.Builder
	var second = input[1 : 1+1]
	var i = 0

	for {

		if i > len([]rune(input))-1 {
			i = 0
		}

		if i%2 == 0 {
			i++
			continue
		}

		s := input[i : i+1]

		if s == second {
			break
		}

		result.WriteString(s)
		i++
	}

	return result.String()
}

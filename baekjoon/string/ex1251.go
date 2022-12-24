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

	minCharIdx := getMinCharIndexes(input)

	var result strings.Builder
	result.WriteString(reverseInput1251(input[0 : minCharIdx[0]+1]))
	result.WriteString(reverseInput1251(input[minCharIdx[0] : minCharIdx[1]+1]))
	result.WriteString(reverseInput1251(input[minCharIdx[1]:]))
	fmt.Println(result.String())
}

func getMinCharIndexes(input string) []int {

	ln := len([]rune(input))

	var fMin = input[0:1]
	var fIdx = 0

	for i := 1; i < ln-2; i++ {
		if input[i:i+1] < fMin {
			fMin = input[i : i+1]
			fIdx = i
		}
	}

	var sMin = input[fIdx+1 : fIdx+2]
	var sIdx = fIdx + 1

	for i := fIdx + 1; i < ln-1; i++ {
		if input[i:i+1] < sMin {
			sMin = input[i : i+1]
			sIdx = i
		}
	}

	return []int{fIdx, sIdx}
}

func reverseInput1251(input string) string {
	var result strings.Builder

	for i := len([]rune(input)) - 1; i >= 0; i-- {
		result.WriteString(input[i : i+1])
	}

	return result.String()
}

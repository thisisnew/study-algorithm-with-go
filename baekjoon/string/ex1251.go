package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {

	var input string
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &input)

	inputRunes := []rune(input)

	sort.Slice(inputRunes, func(i, j int) bool {
		return inputRunes[i] < inputRunes[j]
	})

	minCharIdx := getMinCharIndexes(inputRunes, input)

	var result strings.Builder
	result.WriteString(reverseInput1251(input[0:minCharIdx[0]]))
	result.WriteString(reverseInput1251(input[minCharIdx[0]+1 : minCharIdx[1]]))
	result.WriteString(reverseInput1251(input[minCharIdx[1]+1:]))
	fmt.Println(result.String())
}

func getMinCharIndexes(inputRunes []rune, input string) []int {

	ln := len(inputRunes)

	var index []int
	var remain = 2
	var fIdx = -1

	for i, r := range inputRunes {
		if len(index) == 2 {
			return index
		}

		for j, in := range input {

			if r == in {
				if j > ln-remain {
					break
				}

				if j == fIdx {
					continue
				}

				index = append(index, i)

				if fIdx == -1 {
					fIdx = i
				}

				remain--
			}
		}
	}

	return index
}

func reverseInput1251(input string) string {

	var result strings.Builder

	for i := len([]rune(input)) - 1; i >= 0; i-- {
		result.WriteString(input[i : i+1])
	}

	return result.String()
}

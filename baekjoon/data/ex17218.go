package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var fInput string
var sInput string
var fSlice []string
var sSlice []string

func main() {

	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &fInput)
	fmt.Fscanln(read, &sInput)

	if len([]rune(sInput)) > len([]rune(fInput)) {
		fInput, sInput = sInput, fInput
	}

	var result strings.Builder

	for i, s := range sInput {

		if !isContainRune(string(s), i) {
			continue
		}

		sSlice = append(sSlice, string(s))
	}

	fmt.Println(result.String())
}

func isContainRune(s string, idx int) bool {

	for i := idx; i < len([]rune(fInput)); i++ {
		if s == fInput[i:i+1] {
			fSlice = append(fSlice, fInput[i:i+1])
			return true
		}
	}

	return false
}

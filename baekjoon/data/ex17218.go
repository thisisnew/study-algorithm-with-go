package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var shortInput string
var longInput string
var result string

func main() {

	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &shortInput)
	fmt.Fscanln(read, &longInput)

	if len([]rune(shortInput)) > len([]rune(longInput)) {
		longInput, shortInput = shortInput, longInput
	}

	var sIdx int

	for {
		if sIdx == len([]rune(shortInput))-1 {
			break
		}

		temp := getCommonPassWord(sIdx)

		if len([]rune(temp)) > len([]rune(result)) {
			result = temp
		}

		sIdx++
	}

	fmt.Println(result)
}

func getCommonPassWord(sIdx int) string {
	var result strings.Builder
	var lIdx int
	for i := sIdx; i < len([]rune(shortInput)); i++ {

		s := shortInput[i : i+1]

		isNext, idx := isNextCharacterInLongInput(lIdx, s)

		lIdx = idx

		if !isNext {
			continue
		}

		result.WriteString(s)
	}

	return result.String()
}

func isNextCharacterInLongInput(lIdx int, s string) (bool, int) {

	for i := lIdx; i < len([]rune(longInput)); i++ {
		li := longInput[i : i+1]
		if li == s {
			return true, i
		}
	}

	return false, lIdx
}

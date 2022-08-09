package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var fInput string
var sInput string

func main() {

	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &fInput)
	fmt.Fscanln(read, &sInput)

	if len([]rune(sInput)) > len([]rune(fInput)) {
		fInput, sInput = sInput, fInput
	}

	var commonSlice []string

	for _, s := range sInput {

		if !isContainRune(string(s)) {
			continue
		}

		commonSlice = append(commonSlice, string(s))
	}

	fmt.Println(commonSlice)

	var result strings.Builder

	for _, c := range commonSlice {

		isLocatedNext := isLocatedNext(c, commonSlice)

		if isLocatedNext {
			result.WriteString(c)
		}
	}

	fmt.Println(result.String())

}

func isContainRune(s string) bool {

	for i := 0; i < len([]rune(fInput)); i++ {
		if s == fInput[i:i+1] {
			return true
		}
	}

	return false
}

func isLocatedNext(s string, commonSlice []string) bool {

	for _, si := range fInput {

		if !isCharacterContainsCommonSlice(commonSlice, string(si)) {
			continue
		}

		if string(si) != s {
			return false
		}

		return true
	}

	return true
}

func isCharacterContainsCommonSlice(commonSlice []string, s string) bool {
	for _, c := range commonSlice {
		if c == s {
			return true
		}
	}

	return false
}

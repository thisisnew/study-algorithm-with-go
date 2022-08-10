package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var shortInput string
var longInput string

func main() {

	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &shortInput)
	fmt.Fscanln(read, &longInput)

	if len([]rune(shortInput)) > len([]rune(longInput)) {
		longInput, shortInput = shortInput, longInput
	}

	var commonSlice []string

	for _, s := range shortInput {

		if !isContainsInLongInput(s) {
			continue
		}

		commonSlice = append(commonSlice, string(s))
	}

	fmt.Println(commonSlice)

	var result strings.Builder

	for _, li := range longInput {

		ss := string(li)

		if !isCharacterContainsCommonSlice(commonSlice, ss) {
			continue
		}

		if ss != commonSlice[0] {
			commonSlice = commonSlice[1:]
			fmt.Println(commonSlice)
			continue
		}

		result.WriteString(ss)
		commonSlice = commonSlice[1:]
		fmt.Println(commonSlice)
	}

	fmt.Println(result.String())
}

func isContainsInLongInput(s int32) bool {

	for _, l := range longInput {
		if l == s {
			return true
		}
	}

	return false
}

func isCharacterContainsCommonSlice(commonSlice []string, s string) bool {
	for _, c := range commonSlice {
		if c == s {
			return true
		}
	}

	return false
}

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

	for i := 0; i < len([]rune(longInput))-1; i++ {
		li := longInput[i : i+1]

		if !isInShortInput(li) {
			continue
		}

		for j := i + 1; j < len([]rune(longInput)); j++ {
			temp := fmt.Sprintf("%s%s", li, longInput[j:j+1])
			replacedString := replaceCharactersExceptTemp(temp)

			if !isContainWordsInReplacedString(temp, replacedString) {
				continue
			}

			li += longInput[j : j+1]
		}

		if len([]rune(li)) > len([]rune(result)) {
			result = li
		}
	}

	fmt.Println(result)
}

func isInShortInput(li string) bool {

	for i := 0; i < len([]rune(shortInput)); i++ {

		si := shortInput[i : i+1]

		if si == li {
			return true
		}

	}

	return false
}

func replaceCharactersExceptTemp(temp string) string {
	var result strings.Builder

	for i := 0; i < len([]rune(shortInput)); i++ {

		si := shortInput[i : i+1]

		if !strings.Contains(temp, si) {
			continue
		}

		result.WriteString(si)
	}

	return result.String()
}

func isContainWordsInReplacedString(temp, replacedString string) bool {

	if len(strings.TrimSpace(replacedString)) == 0 {
		return false
	}

	for i := 0; i < len([]rune(replacedString)); i++ {

		var rs = replacedString[i : i+1]

		for j := i + 1; j < len([]rune(replacedString)); j++ {
			if !strings.Contains(temp, fmt.Sprintf("%s%s", rs, replacedString[j:j+1])) {
				continue
			}

			rs += replacedString[j : j+1]

			if rs == temp {
				return true
			}
		}
	}

	return false
}

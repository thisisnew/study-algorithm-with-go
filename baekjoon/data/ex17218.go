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

	var temp strings.Builder
	var sIdx int

	for {

		if sIdx == len([]rune(shortInput))-1 {
			break
		}

		sIdx++
	}

}

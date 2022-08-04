package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var fInput string
	var sInput string
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &fInput)
	fmt.Fscanln(read, &sInput)

	if len([]rune(sInput)) > len([]rune(fInput)) {
		fInput, sInput = sInput, fInput
	}

	var result strings.Builder

	for _, s := range sInput {

		if !strings.ContainsRune(fInput, s) {
			continue
		}

		result.WriteRune(s)
	}

	fmt.Println(result.String())
}

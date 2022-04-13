package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var read = bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(read, &n)

	for i := 0; i < n; i++ {
		var a, b string
		fmt.Fscanln(read, &a, &b)

		var answer strings.Builder
		answer.WriteString(a + " & " + b + " are ")

		aRunes := []rune(a)
		sort.Slice(aRunes, func(i int, j int) bool { return aRunes[i] < aRunes[j] })
		bRunes := []rune(b)
		sort.Slice(bRunes, func(i int, j int) bool { return bRunes[i] < bRunes[j] })

		if string(aRunes) != string(bRunes) {
			answer.WriteString("NOT ")
		}

		answer.WriteString("anagrams.")

		fmt.Println(answer.String())
	}

}

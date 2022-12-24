package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var read = bufio.NewReader(os.Stdin)
	t, _, _ := read.ReadLine()
	text := strings.Split(string(t), " ")

	a := text[0]
	b := text[1]

	if len([]rune(a)) > len([]rune(b)) {
		a, b = b, a
	}

	aIdx, bIdx := getIndexBetweenTwoWords(a, b)

	for i := 0; i < len([]rune(b)); i++ {
		for j := 0; j < len([]rune(a)); j++ {
			if i == bIdx {
				fmt.Print(a[j : j+1])
				continue
			}

			if j != aIdx {
				fmt.Print(".")
			} else {
				fmt.Print(b[i : i+1])
			}
		}

		if i < len([]rune(b))-1 {
			fmt.Println()
		}
	}

}

func getIndexBetweenTwoWords(a, b string) (int, int) {
	var aIdx int
	var bIdx int
	var isAssign bool

	for i, ar := range a {
		if isAssign {
			return aIdx, bIdx
		}

		for j, br := range b {
			if ar == br {
				aIdx = i
				bIdx = j
				isAssign = true
				break
			}
		}
	}

	return aIdx, bIdx
}

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

	var aIdx int
	var bIdx int
	for i, ar := range a {
		for j, br := range b {
			if ar == br {
				aIdx = i
				bIdx = j
				break
			}
		}
	}

	for i := 0; i < len([]rune(b)); i++ {
		for j := 0; j < len([]rune(a)); j++ {

			if i == bIdx {
				fmt.Print(a[j : j+1])
			} else {
				if j != aIdx {
					fmt.Print(".")
				} else {
					fmt.Print()
				}
			}

		}
	}

}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var a, b string
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &a, &b)

	var aRunes = []rune(a)
	var bRunes = []rune(b)
	var dif = 0

	for {

		if len(aRunes) == len(bRunes) {
			fmt.Println(dif)
			return
		}

	}

}

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

func getDifBetweenTwoWords(a, b string) int {

	if len([]rune(a)) > len([]rune(b)) {
		a, b = b, a
	}

	var result int

	for i := 0; i < len([]rune(a)); i++ {
		if a[i:i+1] != b[i:i+1] {
			result++
		}
	}

	return result
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var read = bufio.NewReader(os.Stdin)
	var input string
	fmt.Fscanln(read, &input)

	if isHiss(input) {
		fmt.Println("hiss")
	} else {
		fmt.Println("no hiss")
	}
}

func isHiss(input string) bool {

	var sCnt = 0

	for i := 0; i < len([]rune(input)); i++ {
		if input[i:i+1] == "s" {
			sCnt++
		} else {
			sCnt = 0
		}

		if sCnt == 2 {
			return true
		}
	}

	return false
}

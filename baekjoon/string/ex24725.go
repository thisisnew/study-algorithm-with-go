package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var n, m int
	var read = bufio.NewReader(os.Stdin)

	fmt.Fscanln(read, &n, &m)

	var sl = make([][]string, n)

	for i := 0; i < len(sl); i++ {
		var props = make([]string, m)
		var input string
		fmt.Fscanln(read, &input)

		for i, s := range input {
			props[i] = string(s)
		}

		sl[i] = props
	}

	fmt.Println(sl)

}

func readHorizontalWord(sl [][]string) int {

	var result int

	for _, props := range sl {

		var word string
		for _, s := range props {

		}

	}

	return result
}

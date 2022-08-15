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
		var prop = make([]string, m)
		var input string
		fmt.Fscanln(read, &input)

		for i, s := range input {
			prop[i] = string(s)
		}

		sl[i] = prop
	}

	fmt.Println(sl)

}

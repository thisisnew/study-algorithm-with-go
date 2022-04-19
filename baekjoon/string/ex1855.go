package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var n int
	var input string
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &n)
	fmt.Fscanln(read, &input)

	var arr = make([][]string, n)

	for i := range arr {
		arr[i] = make([]string, n)
	}

	var h = 0
	var v = 0
	for i := 0; i < len(input); i++ {

		if i == len(input)-1 {
			break
		}

		if i > 0 && i%n == 0 {
			h++
			v = 0
		}

		arr[h][v] = input[i : i+1]
		v++
	}

	h = 0
	v = 0
	for i := 0; i < len(input); i++ {
		if i > 0 && i%n == 0 {
			v++
			h = 0
		}
		fmt.Print(arr[h][v])
		h++
	}
}

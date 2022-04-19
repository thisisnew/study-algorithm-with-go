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

	var arr = make([][]string, len(input)/n)

	for i := range arr {
		arr[i] = make([]string, n)
	}

	var h = 0
	var v = 0
	var isReverse bool
	for i := 0; i < len(input); i++ {

		if i == len(input)-1 {
			break
		}

		if i > 0 && i%n == 0 {
			h++

			if isReverse {
				v = 0
			} else {
				v = n - 1
			}

			isReverse = !isReverse
		}

		arr[h][v] = input[i : i+1]

		if isReverse {
			v--
		} else {
			v++
		}
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

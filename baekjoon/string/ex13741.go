package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var input string
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &input)

	var index = int('a')
	var result = 0

	for _, c := range input {

		if index != int(c) {
			if index == int('a') {
				continue
			}

			result++
			continue
		}

		index++
	}

	fmt.Println(result)

}

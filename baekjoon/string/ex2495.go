package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var read = bufio.NewReader(os.Stdin)
	var max int
	for i := 0; i < 3; i++ {
		var input string
		fmt.Fscanln(read, &input)
		cnt := countContinuityNumbers(input)
		if max < cnt {
			max = cnt
		}
	}

	fmt.Println(max)

}

func countContinuityNumbers(input string) int {

	var out = 0

	var prev string
	var cnt = 1
	for i := 1; i < len(input); i++ {
		prev = input[i-1 : i]

		if input[i:i+1] == prev {
			continue
		}

		if cnt > out {
			out = cnt
		}

		cnt = 1
	}

	return out

}

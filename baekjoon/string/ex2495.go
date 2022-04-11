package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var read = bufio.NewReader(os.Stdin)

	for i := 0; i < 3; i++ {
		var input string
		fmt.Fscanln(read, &input)
		countContinuityNumbers(input)
	}

}

func countContinuityNumbers(input string) {

	var max = 1
	var cnt = 1
	var prev = input[0:1]

	for i := 1; i < len(input); i++ {

		if input[i:i+1] == prev {
			cnt++
		} else {
			cnt = 1
		}

		if cnt > max {
			max = cnt
		}

		prev = input[i : i+1]
	}

	fmt.Println(max)
}

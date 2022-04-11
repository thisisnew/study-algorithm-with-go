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

	}

}

func countContinuityNumbers(input string) int{

	var cnt = 0

	var prev string
	for i:=1; i<len(input); i++{
		prev = input[i-1:i]

		if


	}


}

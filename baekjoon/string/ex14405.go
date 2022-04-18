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

	var yes = true
	var i int

	for {

		if i >= len(input) {
			break
		}

		if i+1 < len(input) && input[i:i+1] == "p" && input[i+1:i+2] == "i" {
			i += 2
		} else if i+1 < len(input) && input[i:i+1] == "k" && input[i+1:i+2] == "a" {
			i += 2
		} else if i+2 < len(input) && input[i:i+1] == "c" && input[i+1:i+2] == "h" && input[i+2:i+3] == "u" {
			i += 3
		} else {
			yes = false
			break
		}
	}

	if !yes {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}

}

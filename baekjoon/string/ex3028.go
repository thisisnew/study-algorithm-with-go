package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var order string
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &order)

	var cups = []bool{true, false, false}

	for _, o := range order {
		switch o {
		case 'A':
			cups[0], cups[1] = cups[1], cups[0]
		case 'B':
			cups[1], cups[2] = cups[2], cups[1]
		case 'C':
			cups[0], cups[2] = cups[2], cups[0]
		}
	}

	for i, c := range cups {
		if c {
			fmt.Println(i + 1)
			return
		}
	}
}

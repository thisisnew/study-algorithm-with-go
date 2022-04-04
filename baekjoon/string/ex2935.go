package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var a uint
	var op string
	var b uint

	var read = bufio.NewReader(os.Stdin)

	fmt.Fscanln(read, &a)
	fmt.Fscanln(read, &op)
	fmt.Fscanln(read, &b)

	fmt.Println(calc(a, b, op))
}

func calc(a, b uint, op string) uint {

	var sum uint

	switch op {
	case "+":
		sum = a + b
	case "*":
		sum = a * b
	}

	return sum
}

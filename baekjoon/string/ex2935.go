package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {

	var a big.Int
	var op string
	var b big.Int

	var read = bufio.NewReader(os.Stdin)

	fmt.Fscanln(read, &a)
	fmt.Fscanln(read, &op)
	fmt.Fscanln(read, &b)

	fmt.Println(calc(a, b, op))
}

func calc(a, b big.Int, op string) string {

	sum := new(big.Int)

	switch op {
	case "+":
		sum = sum.Add(&a, &b)
	case "*":
		sum = sum.Mul(&a, &b)
	}

	return sum.Text(10)
}

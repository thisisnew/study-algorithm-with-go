package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var read = bufio.NewReader(os.Stdin)
	var n, m int
	var f string
	var s string
	fmt.Fscanln(read, &n, &m)
	fmt.Fscanln(read, &s)
	fmt.Fscanln(read, &f)

	if n < m {
		n, m = m, n
		f, s = s, f
	}

	for i := 0; i < n; i++ {

	}

	var result int

	cal := calculateRune()

	if result > cal {
		result = cal
	}

	fmt.Println(result)
}

func calculateRune(a, b rune) int {
	var result int

	return result
}

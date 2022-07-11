package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var t int
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscanln(read, &n)

		if isSymmetry(n) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}

func isSymmetry(n int) bool {

	total := n + getReversedNumber(n)
	totalReverse := getReversedNumber(total)

	return total == totalReverse
}

func getReversedNumber(n int) int {

	m := strconv.Itoa(n)

	var sb strings.Builder

	for i := len([]rune(m)) - 1; i >= 0; i-- {
		sb.WriteString(m[i : i+1])
	}

	nReverse, _ := strconv.Atoi(sb.String())

	return nReverse
}

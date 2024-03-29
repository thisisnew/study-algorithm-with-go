package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var read = bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscanln(read, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscanln(read, &n)

		text, _, _ := read.ReadLine()
		values := strings.Split(string(text), " ")

		var m = map[string]int{}
		for j := 0; j < n; j++ {
			m[values[j]] = j
		}

		text, _, _ = read.ReadLine()
		values = strings.Split(string(text), " ")

		var keys = make([]int, n)
		for j := 0; j < n; j++ {
			keys[j] = m[values[j]]
		}

		text, _, _ = read.ReadLine()
		values = strings.Split(string(text), " ")

		var result = make([]string, n)

		for j := 0; j < n; j++ {
			result[keys[j]] = values[j]
		}

		for j := 0; j < n; j++ {
			fmt.Print(result[j], " ")
		}

		fmt.Println()
	}
}

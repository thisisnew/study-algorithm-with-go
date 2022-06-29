package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &n)

	text, _, _ := read.ReadLine()

	sl := strings.Split(string(text), " ")

	var result []string

	for _, s := range sl {
		result = append(result, s)
	}

	fmt.Println(len(result))
}

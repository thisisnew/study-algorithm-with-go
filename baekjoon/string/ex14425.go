package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	var m int

	var result = bufio.NewReader(os.Stdin)
	fmt.Fscanln(result, &n, &m)

	words := map[string]string{}

	for i := 0; i < n; i++ {
		var word string
		fmt.Fscanln(result, &word)
		words[word] = word
	}

	var cnt int
	for i := 0; i < m; i++ {
		var word string
		fmt.Fscanln(result, &word)
		if _, ok := words[word]; ok {
			cnt++
		}
	}

	fmt.Println(cnt)
}

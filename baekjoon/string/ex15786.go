package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var read = bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscanln(read, &n, &m)

	var input string
	fmt.Fscanln(read, &input)

	var sl = make([]string, n)

	for i := 0; i < n; i++ {
		sl[i] = input[i : i+1]
	}

	var words = make([]string, m)
	for i := 0; i < m; i++ {
		var word string
		fmt.Fscanln(read, &word)
		words[i] = word
	}

	for _, word := range words {
		printSendMeTheMoney(sl, word)
	}

}

func printSendMeTheMoney(sl []string, word string) {

	var idx = 0
	var count = len(sl)

	for _, w := range word {
		if string(w) == sl[idx] {
			count--
			idx++
		}

		if count == 0 {
			break
		}
	}

	if count <= 0 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var n int
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &n)

	bookMap := make(map[string]int, n)

	for i := 0; i < n; i++ {
		var book string
		fmt.Fscanln(read, &book)
		bookMap[book]++
	}

	for _, book := range bookMap {

	}

}

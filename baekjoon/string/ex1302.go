package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

	var n int
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &n)

	bookMap := make(map[string]int, n)
	var max int

	for i := 0; i < n; i++ {
		var book string
		fmt.Fscanln(read, &book)
		bookMap[book]++

		if bookMap[book] > max {
			max = bookMap[book]
		}
	}

	var books []string

	for k, v := range bookMap {
		if v != max {
			continue
		}

		books = append(books, k)
	}

	if len(books) > 1 {
		sort.Slice(books, func(i, j int) bool {
			return books[i] < books[j]
		})
	}

	fmt.Println(books[0])
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var read = bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(read, &n)

	var seats = map[string]bool{}

	text, _, _ := read.ReadLine()
	reqs := strings.Split(string(text), " ")

	var result int

	for _, seat := range reqs {

		_, ok := seats[seat]

		if ok {
			result++
			continue
		}

		seats[seat] = true
	}

	fmt.Println(result)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var read = bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(read, &n)

	var seats = make([]int, n+1)

	text, _, _ := read.ReadLine()
	reqs := strings.Split(string(text), " ")

	var result int

	for _, seat := range reqs {

		n, _ := strconv.Atoi(seat)

		if seats[n] != 0 {
			result++
			continue
		}

		seats[n]++
	}

	fmt.Println(result)
}

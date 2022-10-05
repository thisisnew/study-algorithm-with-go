package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {

	var read = bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(read, &n)

	input, _, _ := read.ReadLine()
	inputs := strings.Split(string(input), " ")

	for i, num := range inputs {

	}

}

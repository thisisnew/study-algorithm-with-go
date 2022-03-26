package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var input string
	var read = bufio.NewReader(os.Stdin)

	fmt.Fscan(read, &input)

	var suffixArray []string

	for i, _ := range input {
		suffixArray = append(suffixArray, input[i:])
	}

	sort.Slice(suffixArray, func(i, j int) bool {
		return suffixArray[i] < suffixArray[j]
	})

	fmt.Println(suffixArray)
}

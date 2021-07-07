package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var input string
	var reader = bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &input)

	var alphabet = make(map[int]int)
	for i := 0; i < 26; i++ {
		alphabet[i+97] = -1
	}

	for i := 0; i < len(input); i++ {
		ch, _ := strconv.Atoi(fmt.Sprintf("%d", input[i]))
		if alphabet[ch] == -1 {
			alphabet[ch] = i
		}
	}

	for i := 0; i < 26; i++ {
		fmt.Printf("%d ", alphabet[i+97])
	}
	fmt.Print("\n")
}

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

	for i := 0; i < n; i++ {
		text, _, _ := read.ReadLine()

		inputs := strings.Split(string(text), " ")
		input := strings.Join(inputs, "")

		calculateWord(input)
	}
}

func calculateWord(input string) {

	var result int32

	for _, ch := range input {
		result += ch - 65 + 1
	}

	if result == 100 {
		fmt.Println("PERFECT LIFE")
	} else {
		fmt.Println(result)
	}

}

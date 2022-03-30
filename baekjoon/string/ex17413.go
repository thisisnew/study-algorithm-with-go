package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var input string
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscan(read, &input)

	inputSplit := strings.Split(input, " ")

	var result []string
	for _, item := range inputSplit {
		result = append(result, reverseInput(item))
	}

	fmt.Println(strings.Join(result, " "))

}

func reverseInput(input string) string {
	var reverse string

	for i := len(input) - 1; i >= 0; i-- {
		reverse += input[i : i+1]
	}

	return reverse
}

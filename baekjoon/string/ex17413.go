package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	var input string
	var read = bufio.NewScanner(os.Stdin)

	isRead := read.Scan()

	if isRead {
		input = read.Text()

		if err := read.Err(); err != nil {
			log.Fatalf("Error while reading file: %s", err)
		}

		inputSplit := strings.Split(input, " ")

		var result []string
		for _, item := range inputSplit {
			result = append(result, reverseInput(item))
		}

		fmt.Println(strings.Join(result, " "))
	}

}

func reverseInput(input string) string {
	var reverse string

	for i := len(input) - 1; i >= 0; i-- {
		reverse += input[i : i+1]
	}

	return reverse
}

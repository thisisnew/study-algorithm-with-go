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

	input, _ = read.ReadString('\n')

	inputSplit := strings.Split(input, " ")

	var resultArr []string
	for _, item := range inputSplit {

		isSymbol := false
		for i := 0; i < len(item); i++ {

			runes := []rune(item)
			var result string
			var reversed string
			for _, rn := range runes {
				if rn == '<' {
					isSymbol = true
					result += reverseInput(reversed)
					reversed = ""
				}

				if isSymbol {
					result += string(rn)

					if rn == '>' {
						isSymbol = false
					}
				} else {
					reversed += string(rn)
				}
			}

			resultArr = append(resultArr, result)
		}
	}

	fmt.Print(strings.Join(resultArr, " "))
}

func reverseInput(input string) string {
	var reverse string

	for i := len(input) - 1; i >= 0; i-- {
		reverse += input[i : i+1]
	}

	return reverse
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var input string
	var read = bufio.NewReader(os.Stdin)

	fmt.Fscan(read, &input)

	items := spiltAsMinusSymbol(input)

	fmt.Println(calculate(items))
}

func calculate(items []string) int {

	var total int

	for i, item := range items {
		props := spiltAsPlusSymbol(item)

		if i == 0 {
			total += sum(props)
		} else {
			total -= sum(props)
		}
	}

	return total
}

func sum(props []string) int {

	var sum int

	for _, prop := range props {
		num, _ := strconv.Atoi(prop)

		sum += num
	}

	return sum
}

func spiltAsPlusSymbol(input string) []string {
	return strings.Split(input, "+")
}

func spiltAsMinusSymbol(input string) []string {
	return strings.Split(input, "-")
}

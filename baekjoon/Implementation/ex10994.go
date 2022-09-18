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

	if n == 1 {
		fmt.Println("*")
		return
	}

	var max = 4*(n-1) + 1
	for i := 0; i < max; i++ {
		switch {
		case i == 0, i == max-1:
			printingStarsByNums(max)
		case i%2 == 0:
			printingStarsByNumsEndToEnd(max)
		default:

		}
	}

}

func printingStarsByNumsEndToEnd(num int) {

	var result strings.Builder

	for i := 0; i < num; i++ {

		if i == 0 || i == num-1 {
			result.WriteString("*")
			continue
		}
		result.WriteString(" ")
	}

	fmt.Println(result.String())
}

func printingStarsByNums(num int) {

	var result strings.Builder

	for i := 0; i < num; i++ {
		result.WriteString("*")
	}

	fmt.Println(result.String())
}

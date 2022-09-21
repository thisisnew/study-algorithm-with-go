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

	var loops = 4*(n-1) + 1
	var stars = loops - 2

	for i := 0; i < loops; i++ {
		var result strings.Builder

		for j := 0; j < loops; j++ {

			if i == 0 || i == loops-1 {
				result.WriteString("*")
			}

			if i%2 == 0 {

				continue
			}

			result = printingStarsByNumsEndToEnd(stars, result)
		}

		fmt.Println(result.String())
		result.Reset()
	}

}

func printingStarsByNumsByRules10994(max, mid int) {

	var result strings.Builder

	for i := 0; i < max; i++ {

		if i == 0 || i == max-1 {
			result.WriteString("*")
			continue
		}

		if i == 1 || i == max-2 {
			result.WriteString(" ")
			continue
		}

		result.WriteString("*")
	}

	fmt.Println(result.String())
}

func printingStarsByNumsEndToEnd(stars int, result strings.Builder) strings.Builder {

	for i := 0; i < stars; i++ {

		if i == 0 || i == stars-1 {
			result.WriteString("*")
			continue
		}

		result.WriteString(" ")
	}

	return result
}

func printingStarsByNums(num int) {

	var result strings.Builder

	for i := 0; i < num; i++ {
		result.WriteString("*")
	}

	fmt.Println(result.String())
}

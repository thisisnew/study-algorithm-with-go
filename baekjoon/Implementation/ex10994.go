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
				result = printingStarsByNums(stars, result)
				continue
			}

			result = printingStarsByNumsEndToEnd(stars, result)
		}

		fmt.Println(result.String())
		result.Reset()
	}

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

func printingStarsByNums(stars int, result strings.Builder) strings.Builder {

	for i := 0; i < stars; i++ {
		result.WriteString("*")
	}

	return result
}

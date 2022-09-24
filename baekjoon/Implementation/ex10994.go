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
		dif := loops - stars

		for j := 0; j < loops; j++ {

			if dif > 0 {

				dif -= 2
				continue
			}

			if j == 0 || j == loops-1 {
				result.WriteString("*")
			} else if i%2 == 0 {
				result = printingStarsByNums(stars, result)
			} else {
				result = printingStarsByNumsEndToEnd(stars, result)
			}

			stars -= 2
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

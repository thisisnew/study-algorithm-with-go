package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &n)

	var loops = 4*(n-1) + 1
	var stars = make([][]string, loops)

	for i := 0; i < loops; i++ {
		stars[i] = make([]string, loops)
		for j := 0; j < loops; j++ {
			stars[i][j] = " "
		}
	}

	addStars(0, loops, stars)

	for i := 0; i < loops; i++ {
		for j := 0; j < loops; j++ {
			fmt.Print(stars[i][j])
		}
		fmt.Println()
	}
}

func addStars(n, loops int, stars [][]string) {

	for i := n; i < loops; i++ {
		stars[n][i] = "*"
		stars[loops-1][i] = "*"
		stars[i][n] = "*"
		stars[i][loops-1] = "*"
	}

	if loops == 1 {
		return
	}

	addStars(n+2, loops-2, stars)
}

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

		switch {
		case i%2 == 0:
			addStarsByNumsToBuilder(stars, loops)
			if i < loops/2 {
				stars -= 2
			} else {
				//반대방향
			}
		case i%2 != 0:
			fmt.Println()
			//addStarsByNumsEndToEndToBuilder(stars, loops)
		}

	}

}

func addStarsByNumsEndToEndToBuilder(stars, loops int) {

	var result strings.Builder

	for i := 0; i < stars; i++ {

		if i == 0 || i == stars-1 {
			result.WriteString("*")
			continue
		}

		result.WriteString(" ")
	}

	fmt.Println(result.String())
}

func addStarsByNumsToBuilder(stars, loops int) {

	var result strings.Builder
	dif := loops - stars
	subDif := dif / 2
	starCnt := 0
	subDifCnt := 0
	isBlank := false

	for i := 0; i < loops; i++ {

		if subDifCnt <= subDif {

			if isBlank {
				result.WriteString(" ")
			} else {
				result.WriteString("*")
			}

			isBlank = !isBlank
			subDifCnt++
		} else {
			result.WriteString("*")
			starCnt++
		}

		if subDifCnt >= subDif && starCnt == stars {
			starCnt = 0
			subDifCnt = 1
			isBlank = !isBlank
		}
	}

	fmt.Println(result.String())
}

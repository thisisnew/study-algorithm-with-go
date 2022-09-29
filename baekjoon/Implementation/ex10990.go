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

	var loops = 2*(n-1) + 1
	var mid = loops / 2
	var left = mid - 1
	var right = mid + 1

	for i := 0; i < n; i++ {
		var result strings.Builder

		for j := 0; j < loops; j++ {
			if i == 0 {
				if j == mid {
					result.WriteString("*")
					continue
				}

				result.WriteString(" ")
			} else {
				if j == left || j == right {
					result.WriteString("*")
					continue
				}

				result.WriteString(" ")
			}
		}

		left--
		right++
		fmt.Println(result.String())
		result.Reset()
	}

}

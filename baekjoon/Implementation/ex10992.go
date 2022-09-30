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
	var left = mid
	var right = mid

	for i := 0; i < n; i++ {
		var result strings.Builder

	lp:
		for j := 0; j < loops; j++ {

			switch {
			case i == 0:
				if j == mid {
					result.WriteString("*")
					break lp
				}

				result.WriteString(" ")
			case i == n-1:
				result.WriteString("*")
			default:
				if j == left {
					result.WriteString("*")
					continue
				}

				if j == right {
					result.WriteString("*")
					break lp
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

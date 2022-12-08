package main

import (
	"fmt"
)

func main() {
	fmt.Println(점찍기(1, 5))
}

func 점찍기(k int, d int) int64 {

	var result int64
	var x int64
	var y int64
	var dsq = int64(d) * int64(d)

	for {
		if y*y > dsq {
			if x*x > dsq {
				return result
			}

			x += int64(k)
			y = 0
		}

		if (x*x)+(y*y) <= dsq {
			result++
		}

		y += int64(k)
	}
}

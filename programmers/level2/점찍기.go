package main

import (
	"fmt"
)

func main() {
	fmt.Println(점찍기(1, 5))
}

func 점찍기(k int, d int) int64 {

	var result int64
	var x int
	var y int

	for {
		if x*x > d*d || y*y > d*d {
			return result
		}

		if (x*x)+(y*y) <= d*d {
			result++
			y += k
		}

		x += k
		y = 0
	}
}

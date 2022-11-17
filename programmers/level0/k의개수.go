package main

import "fmt"

func main() {
	fmt.Println(k의개수(1, 13, 1))
}

func k의개수(i int, j int, k int) int {
	return countDivisions(i, j, k)
}

func countDivisions(i int, j int, k int) int {

	var result int

	for a := i; a <= j; a++ {

		var n = a
		var share int
		var reminder int

		for {
			share = n / 10
			reminder = n % 10

			if reminder == k {
				result++
			}

			if share == 0 {
				break
			}

			n = share
		}
	}

	return result
}

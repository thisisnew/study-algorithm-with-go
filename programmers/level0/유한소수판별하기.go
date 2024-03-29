package main

import "fmt"

func main() {
	fmt.Println(유한소수판별하기(11, 22))
}

func 유한소수판별하기(a int, b int) int {

	for {
		d := getDivisionsBetweenNumbers(a, b)

		if d > 1 {
			a = a / d
			b = b / d
			continue
		}

		if isFiniteNumber(b) {
			return 1
		}

		return 2
	}
}

func isFiniteNumber(b int) bool {

	for {
		if b%2 == 0 {
			b /= 2
			continue
		}

		if b%5 == 0 {
			b /= 5
			continue
		}

		return b == 1
	}
}

func getDivisionsBetweenNumbers(a, b int) int {
	if b == 0 {
		return a
	} else {
		return getDivisionsBetweenNumbers(b, a%b)
	}
}

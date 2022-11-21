package main

import "fmt"

func main() {
	fmt.Println(유한소수판별하기(7, 20))
}

func 유한소수판별하기(a int, b int) int {

	if b%a == 0 {
		b = b / a
	}

	for {
		d := getDivisionsBetweenNumbers(a, b)

		if d == a {

			if isFiniteNumber(b) {
				return 1
			}

			return 2

		}

		a = a / d
		b = b / d
	}
}

func isFiniteNumber(b int) bool {

	for {

		if b%2 == 0 {
			b /= 2
		}

		if b%5 == 0 {
			b /= 5
		}

		if b == 2 || b == 5 {
			continue
		}

		if b == 1 {
			return true
		}

		if b != 1 {
			return false
		}
	}
}

func getDivisionsBetweenNumbers(a, b int) int {

	if b == 0 {
		return a
	} else {
		return gcdPoints(b, a%b)
	}

}

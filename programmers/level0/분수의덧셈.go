package main

import (
	"fmt"
)

func main() {
	fmt.Println(분수의덧셈(1, 2, 3, 4))
}

func 분수의덧셈(denum1 int, num1 int, denum2 int, num2 int) []int {

	denum := getDenum(denum1, denum2)
	num := (num1 * (denum / denum1)) + (num2 * (denum / denum2))

	for {
		result := gcdPoints(num, denum)

		if result == 1 || num/result == 1 || denum/result == 1 {
			break
		}

		num = num / result
		denum = denum / result
	}

	return []int{denum, num}
}

func getDenum(denum1 int, denum2 int) int {
	return lcmPoints(denum1, denum2)
}

func lcmPoints(num1 int, num2 int) int {
	gcd := gcdPoints(num1, num2)
	return num1 * num2 / gcd
}

func gcdPoints(num1 int, num2 int) int {
	if num2 == 0 {
		return num1
	} else {
		return gcdPoints(num2, num1%num2)
	}
}

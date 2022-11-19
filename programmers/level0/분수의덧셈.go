package main

import (
	"fmt"
)

func main() {
	fmt.Println(분수의덧셈(512, 512, 512, 512))
}

func 분수의덧셈(denum1 int, num1 int, denum2 int, num2 int) []int {

	denominator := getDenominator(num1, num2)                                      //분모
	numerator := (denum1 * (denominator / num1)) + (denum2 * (denominator / num2)) //분자

	for {
		result := gcdPoints(denominator, numerator)

		if result > 1 && numerator%result == 0 {
			denominator = denominator / result
			numerator = numerator / result
			continue
		}

		if result == 1 || denominator/result == 1 || numerator/result == 1 {
			break
		}

		denominator = denominator / result
		numerator = numerator / result
	}

	return []int{numerator, denominator}
}

func getDenominator(num1 int, num2 int) int {
	return lcmPoints(num1, num2)
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

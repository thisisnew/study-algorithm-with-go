package main

import "fmt"

func main() {
	fmt.Println(합성수찾기(10))
}

func 합성수찾기(n int) int {

	var result int

	for i := 4; i <= n; i++ {
		if isComplexNumber(i) {
			result++
		}
	}

	return result
}

func isComplexNumber(n int) bool {

	var result int

	for i := 1; i*i <= n; i++ {

		if i*i == n {
			result++
			continue
		}

		if n%i == 0 {
			result += 2
		}

	}

	return result >= 3
}

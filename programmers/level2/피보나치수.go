package main

import "fmt"

func main() {
	fmt.Println(피보나치수(3))
}

func 피보나치수(n int) int {

	var result = fibonacci(n)

	return result % 1234567

}

func fibonacci(n int) int {

	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

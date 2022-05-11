package main

import "fmt"

func main() {
	fmt.Println(피보나치수(3))
}

func 피보나치수(n int) int {

	var result = fibo(n)

	return result % 1234567

}

func fibo(n int) int {

	if n <= 1 {
		return n
	}

	return fibo(n-1) + fibo(n-2)
}

package main

import "fmt"

func main() {
	fmt.Println(피보나치수(5))
}

func 피보나치수(n int) int {

	var fibo []int

	for i := 0; i <= n; i++ {

		switch i {
		case 0:
			fibo = append(fibo, 0)
		case 1:
			fibo = append(fibo, 1)
		default:
			fibo = append(fibo, (fibo[i-1]+fibo[i-2])%1234567)
		}

	}

	return fibo[n]
}

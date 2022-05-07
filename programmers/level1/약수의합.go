package main

import "fmt"

func main() {
	fmt.Println(약수의합(5))
}

func 약수의합(n int) int {

	var result int

	for i := 1; i <= n; i++ {

		if n%i != 0 {
			continue
		}

		result += i
	}

	return result
}

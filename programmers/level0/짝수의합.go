package main

import "fmt"

func main() {
	fmt.Println(짝수의합(10))
}

func 짝수의합(n int) int {

	var result int

	for i := 0; i <= n; i += 2 {
		result += i
	}

	return result
}

package main

import "fmt"

func main() {
	fmt.Println(순서쌍의개수(20))
}

func 순서쌍의개수(n int) int {

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

	return result
}

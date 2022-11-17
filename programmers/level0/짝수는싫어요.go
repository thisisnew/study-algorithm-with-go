package main

import "fmt"

func main() {
	fmt.Println(짝수는싫어요(10))
}

func 짝수는싫어요(n int) []int {

	var result []int

	for i := 1; i <= n; i += 2 {
		result = append(result, i)
	}

	return result
}

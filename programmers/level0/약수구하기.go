package main

import "fmt"

func main() {
	fmt.Println(약수구하기(24))
}

func 약수구하기(n int) []int {

	var result []int

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			result = append(result, i)
		}
	}

	return result
}

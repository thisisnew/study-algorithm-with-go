package main

import "fmt"

func main() {
	fmt.Println(자연수뒤집어배열로만들기(12345))
}

func 자연수뒤집어배열로만들기(n int64) []int {

	var result []int

	for {
		if n == 0 {
			break
		}

		result = append(result, int(n%10))
		n = n / 10
	}

	return result
}

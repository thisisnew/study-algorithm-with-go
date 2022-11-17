package main

import "fmt"

func main() {
	fmt.Println(n의배수고르기(3, []int{4, 5, 6, 7, 8, 9, 10, 11, 12}))
}

func n의배수고르기(n int, numlist []int) []int {

	var result []int

	for _, num := range numlist {
		if num%n == 0 {
			result = append(result, num)
		}
	}

	return result
}

package main

import "fmt"

func main() {
	nums := []int{
		1, 2, 7, 6, 4,
	}

	fmt.Println(소수만들기(nums))
}

func 소수만들기(nums []int) int {
	var result int

	for i := 0; i < len(nums)-2; i++ {

		n := nums[i] + nums[i+1] + nums[i+2]

		if !isPrimeNumber(n) {
			continue
		}

		result++
	}

	return result
}

func isPrimeNumber(n int) bool {

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

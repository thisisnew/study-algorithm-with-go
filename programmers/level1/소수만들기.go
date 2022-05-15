package main

import "fmt"

func main() {
	nums := []int{
		1, 2, 3, 4,
	}

	fmt.Println(소수만들기(nums))
}

func 소수만들기(nums []int) int {
	var result int

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				n := nums[i] + nums[j] + nums[k]

				if isPrimeNumber(n) {
					result++
				}

			}
		}
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

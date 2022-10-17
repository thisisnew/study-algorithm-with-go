package main

import "fmt"

func main() {
	nums := []int{
		-1, 1, -1, 1,
	}

	fmt.Println(삼총사(nums))
}

func 삼총사(nums []int) int {

	var result int

	for i := 0; i < len(nums)-2; i++ {

		for j := i + 1; j < len(nums)-1; j++ {

			for k := j + 1; k < len(nums); k++ {
				r := nums[i] + nums[j] + nums[k]
				if r == 0 {
					result++
				}
			}
		}
	}

	return result
}

package main

import "fmt"

func main() {
	nums := []int{
		-2, 3, 0, 2, -5,
	}

	fmt.Println(삼총사(nums))
}

func 삼총사(nums []int) int {

	var result int

	for i := 0; i < len(nums); i++ {
	se:
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue se
			}
		th:
			for k := 0; k < len(nums); k++ {
				if k == i || k == j {
					continue th
				}

				r := nums[i] + nums[j] + nums[k]
				if r == 0 {
					result++
				}
			}
		}
	}

	return result
}

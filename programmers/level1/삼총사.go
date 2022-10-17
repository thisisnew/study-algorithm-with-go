package main

import "fmt"

func main() {
	nums := []int{
		1, 2, 3, 4,
	}

	fmt.Println(삼총사(nums))
}

func 삼총사(nums []int) int {

	var result int

	for i := 0; i < len(nums); i++ {
		var r = i
	se:
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue se
			}
			r += j
		th:
			for k := 0; k < len(nums); k++ {
				if k == i || k == j {
					continue th
				}

				if r+k == 0 {
					result++
				}
			}
		}
	}

	return result
}

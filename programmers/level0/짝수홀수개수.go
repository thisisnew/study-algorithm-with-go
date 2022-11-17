package main

import "fmt"

func main() {
	fmt.Println(짝수홀수개수([]int{1, 2, 3, 4, 5}))
}

func 짝수홀수개수(num_list []int) []int {

	var even int
	var odd int

	for _, num := range num_list {
		if num%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	return []int{even, odd}
}

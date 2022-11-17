package main

import "fmt"

func main() {
	fmt.Println(배열뒤집기([]int{1, 2, 3, 4, 5}))
}

func 배열뒤집기(num_list []int) []int {

	var result = make([]int, len(num_list))

	for i := 0; i < len(result); i++ {
		result[i] = num_list[len(num_list)-1-i]
	}

	return result
}

package main

import "fmt"

func main() {
	fmt.Println(배열두배만들기([]int{1, 2, 3, 4, 5}))
}

func 배열두배만들기(numbers []int) []int {

	var result = make([]int, len(numbers))

	for i, number := range numbers {
		result[i] = number * 2
	}

	return result
}

package main

import "fmt"

const totalZeroToNine = 45

func main() {
	n := []int{1, 2, 3, 4, 6, 7, 8, 0}
	fmt.Println(없는숫자더하기(n))
}

func 없는숫자더하기(numbers []int) int {
	var result int

	for _, n := range numbers {
		result += n
	}

	return totalZeroToNine - result
}

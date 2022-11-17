package main

import "fmt"

func main() {
	fmt.Println(배열의평균값([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}

func 배열의평균값(numbers []int) float64 {

	var sum float64

	for _, number := range numbers {
		sum += float64(number)
	}

	return sum / float64(len(numbers))
}

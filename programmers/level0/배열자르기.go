package main

import "fmt"

func main() {
	fmt.Println(배열자르기([]int{1, 2, 3, 4, 5}, 1, 3))
}

func 배열자르기(numbers []int, num1 int, num2 int) []int {
	return numbers[num1 : num2+1]
}

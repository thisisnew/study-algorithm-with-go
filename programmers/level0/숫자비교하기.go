package main

import "fmt"

func main() {
	fmt.Println(숫자비교하기(2, 3))
}

func 숫자비교하기(num1 int, num2 int) int {

	if num1 == num2 {
		return 1
	}

	return -1
}

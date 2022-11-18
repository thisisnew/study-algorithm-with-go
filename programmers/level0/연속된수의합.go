package main

import "fmt"

func main() {
	fmt.Println(연속된수의합(4, 14))
}

func 연속된수의합(num int, total int) []int {

	midValue := total / num
	var startNum = midValue - (num / 2)

	if total%num != 0 {
		startNum = midValue - (num / 2) + 1
	}

	var result = make([]int, num)

	for i := 0; i < num; i++ {
		result[i] = startNum + i
	}

	return result
}

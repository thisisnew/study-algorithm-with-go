package main

import "fmt"

func main() {
	fmt.Println(짝수와홀수(4))
}

func 짝수와홀수(num int) string {
	var result string

	switch num % 2 {
	case 0:
		result = "Even"
	default:
		result = "Odd"
	}

	return result
}

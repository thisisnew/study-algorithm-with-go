package main

import "fmt"

func main() {
	fmt.Println(짝수와홀수(4))
}

func 짝수와홀수(num int) string {

	switch num % 2 {
	case 0:
		return "Even"
	case 1:
		return "Odd"
	}

	return ""
}

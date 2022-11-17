package main

import "fmt"

func main() {
	fmt.Println(나머지구하기(3, 2))
}

func 나머지구하기(num1 int, num2 int) int {
	return num1 % num2
}

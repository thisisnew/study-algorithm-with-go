package main

import "fmt"

func main() {
	fmt.Println(피자나눠먹기1(7))
}

func 피자나눠먹기1(n int) int {

	var addition = 0

	if n%7 > 0 {
		addition = 1
	}

	return n/7 + addition
}

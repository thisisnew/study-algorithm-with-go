package main

import "fmt"

func main() {
	fmt.Println(피자나눠먹기3(7, 10))
}

func 피자나눠먹기3(slice int, n int) int {

	var addition = 0

	if n%slice > 0 {
		addition = 1
	}

	return n/slice + addition
}

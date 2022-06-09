package main

import (
	"fmt"
)

func main() {
	fmt.Println(멀리뛰기(4))
}

func 멀리뛰기(n int) int64 {

	var result int64

	var sl = make([]int, n)

	for i := 0; i < n; i++ {
		sl[i] = 1
	}

	return result / 1234567
}

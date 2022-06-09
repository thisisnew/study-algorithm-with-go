package main

import (
	"fmt"
)

func main() {
	fmt.Println(멀리뛰기(4))
}

func 멀리뛰기(n int) int64 {

	var sl = make([]int64, n+1)

	for i := 1; i <= n; i++ {

		if i == 1 {
			sl[i] = 1
			continue
		}

		if i == 2 {
			sl[i] = 2
			continue
		}

		sl[i] = (sl[i-2] + sl[i-1]) % 1234567

	}

	return sl[n]
}

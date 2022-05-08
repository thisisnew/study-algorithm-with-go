package main

import "fmt"

func main() {
	fmt.Println(정수제곱근판별(121))
}

func 정수제곱근판별(n int64) int64 {

	var result int64
	var num int64

	for {

		num++

		pow := num * num

		if pow < n {
			continue
		}

		if pow > n {
			result = -1
			break
		}

		if pow == n {
			result = (num + 1) * (num + 1)
			break
		}

	}

	return result
}

package main

import "fmt"

func main() {
	fmt.Println(정수제곱근판별(121))
}

func 정수제곱근판별(n int64) int64 {

	var start int64 = 1
	var result int64

	if n > 100 {
		start = n / 10

		for {
			pow := start * start

			if pow == n {
				result = start + 1
				return result * result
			}

			if pow < n {
				break
			}

			if pow > n {
				start = start / 10
				continue
			}

		}
	}

	for i := start; i < n; i++ {
		pow := i * i

		if pow < n {
			continue
		}

		if pow > n {
			return -1
		}

		if pow == n {
			result = i + 1
			return result * result
		}
	}

	return -1
}

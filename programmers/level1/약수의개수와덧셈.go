package main

import "fmt"

func main() {
	fmt.Println(약수의개수와덧셈(13, 17))
}

func 약수의개수와덧셈(left int, right int) int {

	var result int

	for i := left; i <= right; i++ {

		var cnt int

		for j := 1; j <= i; j++ {

			if i%j != 0 {
				continue
			}

			cnt++
		}

		if cnt%2 == 0 {
			result += i
		} else {
			result -= i
		}

	}

	return result
}

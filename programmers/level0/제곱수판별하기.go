package main

import "fmt"

func main() {
	fmt.Println(제곱수판별하기(144))
}

func 제곱수판별하기(n int) int {

	var x = 1

	for {

		if x*x == n {
			return 1
		}

		if x*x > n {
			return 2
		}

		x++

	}

}

package main

import "fmt"

func main() {
	fmt.Println(유한소수판별하기(7, 20))
}

func 유한소수판별하기(a int, b int) int {

	if b%a == 0 {
		b = b / a
	}

	n := 10

	if 10 < b && b <= 100 {
		n = 100
	}

	if 100 < b && b <= 1000 {
		n = 1000
	}

	if n%b == 0 {
		return 1
	}

	return 2
}

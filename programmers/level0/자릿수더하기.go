package main

import (
	"fmt"
)

func main() {
	fmt.Println(자릿수더하기(1234))
}

func 자릿수더하기(n int) int {

	var result int
	var share int
	var reminder int

	for {
		share = n / 10
		reminder = n % 10

		result += reminder

		if n == share {
			break
		}

		n = share
	}

	return result
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(저주의숫자3(15))
}

func 저주의숫자3(n int) int {

	var result = 0

	for i := 1; i <= n; i++ {
		result++

		for result%3 == 0 || strings.Contains(strconv.Itoa(result), "3") {
			result++
		}
	}

	return result
}

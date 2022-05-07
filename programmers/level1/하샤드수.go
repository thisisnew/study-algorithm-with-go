package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(하샤드수(11))
}

func 하샤드수(x int) bool {
	return x%sumDigit(x) == 0
}

func sumDigit(x int) int {

	var result int

	numStr := strconv.Itoa(x)

	for _, nm := range numStr {

		n, _ := strconv.Atoi(string(nm))

		if n == 0 {
			continue
		}

		result += n
	}

	return result
}

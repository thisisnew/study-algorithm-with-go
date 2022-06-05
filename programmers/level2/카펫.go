package main

import "fmt"

func main() {
	brown := 10
	yellow := 2

	fmt.Println(카펫(brown, yellow))
}

func 카펫(brown int, yellow int) []int {
	sum := brown + yellow

	var result []int

	for i := 1; i < (brown / 2); i++ {

		if i*2+2 != brown {
			continue
		}

		result = []int{i, sum / i}
		break
	}

	return result
}

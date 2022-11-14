package main

import (
	"fmt"
)

func main() {
	fmt.Println(양궁대회(5, []int{2, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0}))
}

func 양궁대회(n int, info []int) []int {

	countZeros := countZeroPoint(info)

	if countZeros > len(info) && n > countZeros {
		return winOverZeroPoint(n, info)
	}

	var result = make([]int, len(info))

	for i, p := range info {

		v := p + 1
		n -= v

		if n < 0 {
			break
		}

		result[i] = v
	}

	return result
}

func countZeroPoint(info []int) int {

	var result int

	for _, p := range info {
		if p == 0 {
			result++
		}
	}

	return result
}

func winOverZeroPoint(n int, info []int) []int {

	var result = make([]int, len(info))
	var remain = n

	for i, p := range info {

		if p > 0 {
			continue
		}

		result[i]++
		remain--
	}

	if remain > 0 {
		result[len(result)-1] += remain
	}

	return result
}

func getNonZeroPointInfo(info []int) []int {

	var result []int

	for _, p := range info {
		if p == 0 {
			continue
		}

		result = append(result, p)
	}

	return result
}

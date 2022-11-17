package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(숫자찾기(232443, 4))
}

func 숫자찾기(num int, k int) int {

	numStr := strconv.Itoa(num)
	kStr := strconv.Itoa(k)

	for i, ns := range numStr {
		if string(ns) == kStr {
			return i + 1
		}
	}

	return -1
}

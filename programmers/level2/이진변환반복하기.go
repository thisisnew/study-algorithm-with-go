package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(이진변환반복하기("110010101001"))
}

func 이진변환반복하기(s string) []int {

	var cnt int
	var result int

	for {

		r, c := countAndRemoveZeros(s)

		if c == 0 {
			return []int{result, cnt}
		}

		result++
		cnt += c

		s = convertDecimalToBinary(len(r))

	}
}

func countAndRemoveZeros(s string) (string, int) {

	var cnt int
	var result string

	for _, r := range s {

		sr := string(r)

		if sr == "0" {
			cnt++
			continue
		}

		result += sr
	}

	return result, cnt
}

func convertDecimalToBinary(l int) string {
	return strconv.FormatInt(int64(l), 2)
}

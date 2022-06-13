package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(이진변환반복하기("01110"))
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
	var result strings.Builder

	for _, r := range s {

		sr := string(r)

		if sr == "0" {
			cnt++
			continue
		}

		result.WriteString(sr)
	}

	return result.String(), cnt
}

func convertDecimalToBinary(l int) string {
	return strconv.FormatInt(int64(l), 2)
}

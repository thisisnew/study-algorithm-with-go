package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(이진변환반복하기("1111111"))
}

func 이진변환반복하기(s string) []int {

	var cnt int
	var removeZeros int

	for s != "1" {

		r, c := countAndRemoveZeros(s)
		cnt++
		removeZeros += c

		s = convertDecimalToBinary(len(r))

	}

	return []int{cnt, removeZeros}
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

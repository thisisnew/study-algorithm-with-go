package main

import "fmt"

func main() {
	fmt.Println(이진변환반복하기())
}

func 이진변환반복하기(s string) []int {
	return []int{}
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

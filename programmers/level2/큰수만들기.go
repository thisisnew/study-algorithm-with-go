package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(큰수만들기("4177252841", 4))
}

func 큰수만들기(number string, k int) string {

	var ln = len([]rune(number)) - k
	var result string
	numSlice := convertStringToArray(number)

	sort.Slice(numSlice, func(i, j int) bool {
		return numSlice[i] > numSlice[j]
	})

	//남은길이체크 + 최대값찾기
	//문자열 재할당당

	for _, num := range numSlice {
		idx := getIndexByToken(num, number)

		if idx < 0 {
			continue
		}

		isLongerThanAnswer := isRemainNumberLongerThanAnswer(idx, number, ln)

		if !isLongerThanAnswer {
			continue
		}

		result += number[idx : idx+1]
		number = number[idx+1:]
		ln--
	}

	return result
}

func convertStringToArray(number string) []string {
	return strings.Split(number, "")
}

func getIndexByToken(num string, number string) int {

	for i, n := range number {

		s := string(n)

		if s == num {
			return i
		}

	}

	return -1
}

func isRemainNumberLongerThanAnswer(idx int, number string, ln int) bool {
	return len([]rune(number[idx:])) > ln
}

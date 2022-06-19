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
	numSlice := convertStringToSlice(number)

	sort.Slice(numSlice, func(i, j int) bool {
		return numSlice[i] > numSlice[j]
	})

	if numSlice[0] == numSlice[len(numSlice)-1] {
		result = number[0:ln]
		return result
	}

out:
	for {
		if ln == 0 {
			break out
		}

		for i, num := range numSlice {

			idx := getIndexByToken(num, number)

			if idx < 0 {
				continue
			}

			isLonger := isRemainNumberLongerThanAnswer(idx, number, ln)

			if !isLonger {
				continue
			}

			ln--
			number = number[idx+1:]
			numSlice = regenerateStringSlice(numSlice, i)
			result += num
			continue out
		}
	}

	return result
}

func convertStringToSlice(number string) []string {
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
	return len([]rune(number[idx:])) >= ln
}

func regenerateStringSlice(numSlice []string, idx int) []string {

	var result []string

	for i, num := range numSlice {
		if i == idx {
			continue
		}

		result = append(result, num)
	}

	return result
}

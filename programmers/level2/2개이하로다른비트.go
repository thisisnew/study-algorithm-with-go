package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	numbers := []int64{2, 7}

	fmt.Println(두개이하로다른비트(numbers))
}

func 두개이하로다른비트(numbers []int64) []int64 {

	var result = make([]int64, len(numbers))

	for i, n := range numbers {

		if n%2 != 0 {
			result[i] = getMinimalNumberHasDifBits(n)
		} else {
			result[i] = n + 1
		}

	}

	return result
}

func getMinimalNumberHasDifBits(n int64) int64 {

	var sb strings.Builder
	var binN = strconv.FormatInt(n, 2)

	if !strings.Contains(binN, "0") {

		sb.WriteString("10")
		sb.WriteString(strings.ReplaceAll(binN[1:], "0", "1"))

	} else {

		lastZeroIdx := strings.LastIndex(binN, "0")
		nextLastZeroIdx := strings.Index(binN[lastZeroIdx:], "1")

		sb.WriteString(binN[0:lastZeroIdx])
		sb.WriteString("10")
		sb.WriteString(binN[nextLastZeroIdx+1:])
	}

	dec, _ := strconv.ParseInt(sb.String(), 2, 64)

	return dec
}

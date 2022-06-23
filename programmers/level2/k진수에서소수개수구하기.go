package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(k진수에서소수개수구하기(437674, 3))
}

func k진수에서소수개수구하기(n int, k int) int {

	var result int

	kString := convertDecimalToKNumber(n, k)

	//kNumber, _ := strconv.Atoi(kString)

	var sb strings.Builder

	for _, num := range kString {

		if string(num) == "0" {
			result += addResultIfPrimeNumber(sb.String())
			sb.Reset()
			continue
		}

		sb.WriteString(string(num))
	}

	return result
}

func convertDecimalToKNumber(n int, k int) string {
	return strconv.FormatInt(int64(n), k)
}

func addResultIfPrimeNumber(num string) int {

	n, _ := strconv.Atoi(num)

	for i := 2; i <= n; i++ {

		if n%i == 0 {
			return 0
		}

	}

	return 1
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(k진수에서소수개수구하기(437674, 3))
}

func k진수에서소수개수구하기(n int, k int) int {

	var result int

	kNumber := convertDecimalToKNumber(n, k)

	return result
}

func convertDecimalToKNumber(n int, k int) string {
	return strconv.FormatInt(int64(n), k)
}

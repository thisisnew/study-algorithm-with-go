package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(이진수더하기("1001", "1111"))
}

func 이진수더하기(bin1 string, bin2 string) string {
	total := getDecimalValue(bin1) + getDecimalValue(bin2)
	return strconv.FormatInt(int64(total), 2)
}

func getDecimalValue(bin string) float64 {

	n, _ := strconv.Atoi(bin)
	var share int
	var reminder int
	var idx float64 = 0
	var result float64

	for {
		share = n / 10
		reminder = n % 10
		result += math.Pow(2.0, idx) * float64(reminder)
		if share == 0 {
			return result
		}
		n = share
		idx++
	}
}

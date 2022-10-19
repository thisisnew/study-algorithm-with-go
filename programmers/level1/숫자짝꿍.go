package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(숫자짝꿍("5525", "1255"))
}

func 숫자짝꿍(X string, Y string) string {

	xNum, _ := strconv.Atoi(X)
	yNum, _ := strconv.Atoi(Y)

	x := getPairSlice(xNum)
	y := getPairSlice(yNum)

	var sl []string

	for i := 9; i >= 0; i-- {
		idx := strconv.Itoa(i)

		for x[i] > 0 && y[i] > 0 {
			sl = append(sl, idx)
			x[i]--
			y[i]--
		}
	}

	if len(sl) == 0 {
		return "-1"
	}

	if sl[0] == "0" {
		return "0"
	}

	return strings.Join(sl, "")
}

func getPairSlice(num int) []int {

	var result = make([]int, 10)

	for {
		result[num%10]++

		r := num / 10

		if r == 0 {
			return result
		}

		num = r
	}
}

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

	x := getPairSlice(X)
	y := getPairSlice(Y)

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

func getPairSlice(num string) []int {

	var result = make([]int, 10)

	for _, s := range num {
		n, _ := strconv.Atoi(string(s))
		result[n]++
	}

	return result
}

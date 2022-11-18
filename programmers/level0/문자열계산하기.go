package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(문자열계산하기("3 + 4"))
}

func 문자열계산하기(my_string string) int {

	sl := strings.Split(my_string, " ")
	var result, _ = strconv.Atoi(sl[0])
	var isPlus = false

	for i := 1; i < len(sl); i++ {
		switch sl[i] {
		case "+":
			isPlus = true
		case "-":
			isPlus = false
		default:

			n, _ := strconv.Atoi(sl[i])

			if isPlus {
				result += n
			} else {
				result -= n
			}

		}
	}

	return result
}

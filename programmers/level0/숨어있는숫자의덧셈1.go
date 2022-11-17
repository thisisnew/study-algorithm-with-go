package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	fmt.Println(숨어있는숫자의덧셈1("aAb1B2cC34oOp"))
}

func 숨어있는숫자의덧셈1(my_string string) int {

	var result int

	for _, s := range my_string {
		if unicode.IsNumber(s) {
			n, _ := strconv.Atoi(string(s))
			result += n
		}
	}

	return result
}

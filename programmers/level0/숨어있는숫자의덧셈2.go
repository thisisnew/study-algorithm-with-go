package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(숨어있는숫자의덧셈2("aAb1B2cC34oOp"))
}

func 숨어있는숫자의덧셈2(my_string string) int {

	var sb strings.Builder
	var result int

	for _, s := range my_string {

		if !unicode.IsNumber(s) {
			if sb.Len() > 0 {
				n, _ := strconv.Atoi(sb.String())
				result += n
				sb.Reset()
			}

			continue
		}

		sb.WriteRune(s)
	}

	if sb.Len() > 0 {
		n, _ := strconv.Atoi(sb.String())
		result += n
	}

	return result
}

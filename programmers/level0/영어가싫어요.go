package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(영어가싫어요("onetwothreefourfivesixseveneightnine"))
}

func 영어가싫어요(numbers string) int64 {

	var numberMap = map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	var sb strings.Builder
	var resultStr strings.Builder

	for _, n := range numbers {
		sb.WriteRune(n)

		if sb.Len() >= 3 {
			num, ok := numberMap[sb.String()]

			if !ok {
				continue
			}

			resultStr.WriteString(num)
			sb.Reset()
		}
	}

	if sb.Len() > 0 {
		num, ok := numberMap[sb.String()]

		if ok {
			resultStr.WriteString(num)
		}
	}

	result, _ := strconv.Atoi(resultStr.String())

	return int64(result)
}

package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(숫자문자열과영단어("one4seveneight"))
}

func 숫자문자열과영단어(s string) int {

	var sb strings.Builder
	var temp strings.Builder

	for _, ch := range s {

		if ch >= 48 && ch <= 57 {
			sb.WriteRune(ch)
			continue
		}

		temp.WriteRune(ch)
		sNum, err := convertWordToNumber(temp.String())

		if err != nil {
			continue
		}

		sb.WriteString(sNum)
		temp.Reset()
	}

	result, _ := strconv.Atoi(sb.String())

	return result
}

func convertWordToNumber(temp string) (string, error) {

	var n int
	var err error

	switch temp {
	case "zero":
		n = 0
	case "one":
		n = 1
	case "two":
		n = 2
	case "three":
		n = 3
	case "four":
		n = 4
	case "five":
		n = 5
	case "six":
		n = 6
	case "seven":
		n = 7
	case "eight":
		n = 8
	case "nine":
		n = 9
	default:
		err = errors.New("not a number word")
	}

	return strconv.Itoa(n), err
}

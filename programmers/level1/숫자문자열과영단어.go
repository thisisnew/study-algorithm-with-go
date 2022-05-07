package main

import (
	"errors"
	"strconv"
	"strings"
)

func main() {

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
		err, n := convertWordToNumber(temp.String())

		if err != nil {
			continue
		}

		sNum := strconv.Itoa(n)
		sb.WriteString(sNum)
		temp.Reset()
	}

	result, _ := strconv.Atoi(sb.String())

	return result
}

func convertWordToNumber(temp string) (error, int) {

	switch temp {
	case "zero":
		return nil, 0
	case "one":
		return nil, 1
	case "two":
		return nil, 2
	case "three":
		return nil, 3
	case "four":
		return nil, 4
	case "five":
		return nil, 5
	case "six":
		return nil, 6
	case "seven":
		return nil, 7
	case "eight":
		return nil, 8
	case "nine":
		return nil, 9
	}

	return errors.New("not a number word"), 0

}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(문자열다루기기본("1234"))

}

func 문자열다루기기본(s string) bool {

	if len(s) != 4 && len(s) != 6 {
		return false
	}

	_, err := strconv.Atoi(s)

	if err != nil {
		return false
	}

	return true
}

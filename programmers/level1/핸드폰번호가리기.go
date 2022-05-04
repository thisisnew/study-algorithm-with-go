package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(핸드폰번호가리기("027778888"))
}

func 핸드폰번호가리기(phoneNumber string) string {

	n := len(phoneNumber) - 4

	return genMasking(n) + phoneNumber[len(phoneNumber)-4:]
}

func genMasking(n int) string {

	var result strings.Builder

	for i := 0; i < n; i++ {
		result.WriteString("*")
	}

	return result.String()
}

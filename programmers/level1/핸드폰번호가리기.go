package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(핸드폰번호가리기("027778888"))
}

func 핸드폰번호가리기(phoneNumber string) string {
	return strings.Repeat("*", len(phoneNumber)-4) + phoneNumber[len(phoneNumber)-4:]
}

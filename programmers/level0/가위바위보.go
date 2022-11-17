package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(가위바위보("2"))
}

func 가위바위보(rsp string) string {

	var result strings.Builder

	for _, s := range rsp {
		str := string(s)

		switch str {
		case "0":
			result.WriteString("5")
		case "2":
			result.WriteString("0")
		case "5":
			result.WriteString("2")

		}
	}

	return result.String()
}

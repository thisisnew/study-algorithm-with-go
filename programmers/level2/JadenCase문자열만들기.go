package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(JadenCase문자열만들기("for the last week"))
}

func JadenCase문자열만들기(s string) string {

	var result []string

	sSlice := strings.Fields(s)

	for _, st := range sSlice {
		result = append(result, fmt.Sprintf("%s%s", strings.ToUpper(st[0:1]), strings.ToLower(st[1:])))
	}

	return strings.Join(result, " ")

}

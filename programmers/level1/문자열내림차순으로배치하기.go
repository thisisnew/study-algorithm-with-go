package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(문자열내림차순으로배치하기("Zbcdefg"))
}

func 문자열내림차순으로배치하기(s string) string {

	result := strings.Split(s, "")

	sort.Slice(result, func(i, j int) bool {
		return result[i] > result[j]
	})

	return strings.Join(result, "")
}

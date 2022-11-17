package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(문자열정렬하기2("Bcad"))
}

func 문자열정렬하기2(my_string string) string {

	var result strings.Builder

	for _, s := range my_string {
		result.WriteString(strings.ToLower(string(s)))
	}

	sl := strings.Split(result.String(), "")

	sort.Slice(sl, func(i, j int) bool {
		return sl[i] < sl[j]
	})

	return strings.Join(sl, "")
}

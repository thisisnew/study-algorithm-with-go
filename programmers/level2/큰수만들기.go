package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(큰수만들기("1231234", 3))
}

func 큰수만들기(number string, k int) string {

	var sl []string

	for _, n := range number {
		sl = append(sl, string(n))
	}

	sort.Slice(sl, func(i, j int) bool {
		return sl[i] > sl[j]
	})

	result := strings.Join(sl, "")

	return result[0 : len(result)-k]
}

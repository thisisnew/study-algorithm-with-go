package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(큰수만들기("1924", 2))
}

func 큰수만들기(number string, k int) string {

	var p = number

	sort.Slice(p, func(i, j int) bool {
		return p[i] > p[j]
	})

	max, _ := strconv.Atoi(p[0:1])

	for i := 0; i < len([]rune(number)); i++ {

		if

	}

	return result.String()
}

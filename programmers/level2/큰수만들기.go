package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(큰수만들기("1924", 2))
}

func 큰수만들기(number string, k int) string {

	var result strings.Builder
	var idx = 0

	for i := 0; i < len([]rune(number)); i++ {
		var max = 0
		for j := idx; j <= k+1; j++ {
			num, _ := strconv.Atoi(number[j : j+1])
			if max < num {
				max = num
				idx = j + 1
			}
		}

		result.WriteString(strconv.Itoa(max))
	}

	return result.String()
}

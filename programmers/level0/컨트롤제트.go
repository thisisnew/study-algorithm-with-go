package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(컨트롤제트("1 2 Z 3"))
}

func 컨트롤제트(s string) int {

	sl := strings.Split(s, " ")

	var result int
	var prev []int

	for _, num := range sl {

		if num == "Z" {
			if len(prev) > 0 {
				p := prev[0]
				prev = prev[1:]
				result -= p
			}

			continue
		}

		n, _ := strconv.Atoi(num)

		result += n
		prev = append([]int{n}, prev...)
	}

	return result
}

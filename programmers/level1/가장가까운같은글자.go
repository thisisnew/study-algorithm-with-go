package main

import "fmt"

func main() {
	fmt.Println(가장가까운같은글자("foobar"))
}

func 가장가까운같은글자(s string) []int {

	var result = make([]int, len([]rune(s)))
	var position = map[rune]int{}

	for i, r := range s {
		p, ok := position[r]

		if !ok {
			result[i] = -1
		} else {
			result[i] = i - p
		}

		position[r] = i
	}

	return result
}

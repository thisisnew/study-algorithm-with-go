package main

import "fmt"

func main() {
	fmt.Println(가장가까운같은글자("foobar"))
}

func 가장가까운같은글자(s string) []int {

	var result []int
	var position = map[rune]int{}

	for i, r := range s {
		p, ok := position[r]

		if !ok {
			result = append(result, -1)
		} else {
			result = append(result, i-p)
		}

		position[r] = i
	}

	return result
}

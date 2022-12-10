package main

import "fmt"

func main() {
	fmt.Println(가장가까운같은글자("banana"))
}

func 가장가까운같은글자(s string) []int {

	var result []int
	var position = map[rune]int{}

	for i, r := range s {
		p, ok := position[r]

		if !ok {
			position[r] = i
			result = append(result, -1)
		} else {
			position[r] = i
			result = append(result, i-p)
		}
	}

	return result
}

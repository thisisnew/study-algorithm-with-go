package main

import (
	"fmt"
)

func main() {
	fmt.Println(예상대진표(8, 4, 7))
}

func 예상대진표(n int, a int, b int) int {

	var round = 1

	for n > 1 {

		var teams = make([][]int, n/2)
		var team = make([]int, 2)
		var idx = 0

		for i := 0; i < n; i++ {
			team[i%2] = i + 1

			if i%2 > 0 {
				teams[idx] = team
				idx++
				team = make([]int, 2)
			}
		}

		for _, team := range teams {
			if !isMatching(team, a, b) {
				continue
			}

			return round
		}

		if a%2 == 0 {
			a = a / 2
		} else {
			a = a/2 + 1
		}

		if b%2 == 0 {
			b = b / 2
		} else {
			b = b/2 + 1
		}

		n = n / 2

		round++
	}

	return round
}

func isMatching(team []int, a, b int) bool {

	isA := false
	isB := false

	for _, t := range team {
		if t == a {
			isA = true
		}

		if t == b {
			isB = true
		}
	}

	return isA && isB
}

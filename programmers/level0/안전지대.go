package main

import "fmt"

func main() {
	fmt.Println(안전지대([][]int{{1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}}))
}

func 안전지대(boards [][]int) int {

	var mined [][]int

	for i, board := range boards {
		for j, b := range board {
			if b == 1 {
				mined = append(mined, []int{i, j})
			}
		}
	}

	for _, m := range mined {
		i := m[0]
		j := m[1]

		if i < len(boards)-1 {
			boards[i+1][j] = 1
		}

		if i < len(boards)-1 && j < len(boards[i])-1 {
			boards[i+1][j+1] = 1
		}

		if i < len(boards)-1 && j > 0 {
			boards[i+1][j-1] = 1
		}

		if i > 0 {
			boards[i-1][j] = 1
		}

		if i > 0 && j < len(boards[i])-1 {
			boards[i-1][j+1] = 1
		}

		if i > 0 && j > 0 {
			boards[i-1][j-1] = 1
		}

		if j < len(boards[i])-1 {
			boards[i][j+1] = 1
		}

		if j > 0 {
			boards[i][j-1] = 1
		}
	}

	var result int
	for _, board := range boards {
		for _, b := range board {
			if b == 0 {
				result++
			}
		}
	}

	return result
}

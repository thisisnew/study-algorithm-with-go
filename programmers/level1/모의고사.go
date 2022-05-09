package main

import "fmt"

var one = []int{1, 2, 3, 4, 5}
var two = []int{2, 1, 2, 3, 2, 4, 2, 5}
var three = []int{3, 3, 1, 1, 2, 2, 4, 4, 5, 5}
var mens = [][]int{one, two, three}

func main() {
	answers := []int{1, 3, 2, 4, 2}
	fmt.Println(모의고사(answers))
}

func 모의고사(answers []int) []int {

	var max int
	var result []int

	for i, men := range mens {

		var idx = 0
		var cnt = 0

		for _, as := range answers {

			if as == men[idx] {
				cnt++
			}

			idx++

			if idx == len(men) {
				idx = 0
			}
		}

		if cnt < max {
			continue
		}

		if cnt > 0 && cnt == max {
			result = append(result, i+1)
			continue
		}

		if cnt > max {
			max = cnt
			result = []int{i + 1}
			continue
		}
	}

	return result
}

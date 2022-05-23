package main

import "fmt"

func main() {

	clothes := [][]string{
		{"crowmask", "face"},
		{"bluesunglasses", "face"},
		{"smoky_makeup", "face"},
	}

	fmt.Println(위장(clothes))
}

func 위장(clothes [][]string) int {

	var resultMap = map[string][]string{}

	for _, cloth := range clothes {

		c, ok := resultMap[cloth[1]]

		if !ok {
			resultMap[cloth[1]] = []string{cloth[0]}
			continue
		}

		c = append(c, cloth[0])
		resultMap[cloth[1]] = c
	}

	var multiplying = 1
	var plus = 0

	for _, v := range resultMap {
		multiplying *= len(v)
		plus += len(v)
	}

	if len(resultMap) == 1 {
		return plus
	} else {
		return multiplying + plus
	}

}

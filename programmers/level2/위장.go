package main

import "fmt"

func main() {

	clothes := [][]string{
		{"yellowhat", "headgear"},
		{"bluesunglasses", "eyewear"},
		{"green_turban", "headgear"},
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

	var result = 1

	for _, v := range resultMap {
		result *= len(v) + 1
	}

	return result - 1

}

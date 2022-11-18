package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(소인수분해(17))
}

func 소인수분해(n int) []int {

	var m = map[int]bool{}
	var idx = 2
	var temp = n

	for {

		if temp%idx == 0 {
			temp = temp / idx

			if _, ok := m[idx]; !ok {
				m[idx] = true
			}

			if temp == 1 || idx == n {
				break
			}

			idx = 2
			continue
		}

		idx++
	}

	var result []int
	for k, _ := range m {
		result = append(result, k)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return result
}

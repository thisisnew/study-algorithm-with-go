package main

import (
	"fmt"
	"sort"
)

func main() {
	priorities := []int{2, 1, 3, 2}
	location := 2
	fmt.Println(print(priorities, location))
}

func print(priorities []int, location int) int {

	var cp = make([]int, len(priorities))
	copy(cp, priorities)

	sort.Slice(cp, func(i, j int) bool {
		return cp[i] > cp[j]
	})

	var count int
	var result int

	for {
		p, popSlice := pop(priorities)
		priorities = popSlice

		if p == cp[0] {
			count++
			_, cp = pop(cp)

			if location == 0 {
				result = count
				break
			}

			location--
		} else {

			pushSlice := push(priorities, p)
			priorities = pushSlice

			if location == 0 {
				location = len(priorities) - 1
			} else {
				location--
			}

		}
	}

	return result
}

func pop(sl []int) (int, []int) {

	var result = sl[0]

	cp := make([]int, len(sl)-1)
	copy(cp, sl[1:])

	return result, cp
}

func push(sl []int, n int) []int {

	cp := make([]int, len(sl)+1)
	copy(cp, sl)
	cp[len(cp)-1] = n

	return cp
}

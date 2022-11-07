package main

import (
	"fmt"
)

func main() {
	fmt.Println(햄버거만들기([]int{2, 1, 1, 2, 3, 1, 2, 3, 1}))
}

func 햄버거만들기(ingredient []int) int {

	var result = 0
	var st []int

	for i := 0; i < len(ingredient); i++ {
		st = append(st, ingredient[i])

		if len(st) >= 4 {

			var temp = [4]int{st[len(st)-4], st[len(st)-3], st[len(st)-2], st[len(st)-1]}

			if temp == [4]int{1, 2, 3, 1} {
				result++

				st = st[1:]
				st = st[1:]
				st = st[1:]
				st = st[1:]
			}

		}

	}

	return result
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println(햄버거만들기([]int{1, 3, 2, 1, 2, 1, 3, 1, 2}))
}

func 햄버거만들기(ingredient []int) int {

	var result = 0
	var st = make([]int, len(ingredient))

	for i := 0; i < len(ingredient); i++ {
		st[i] = ingredient[i]

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

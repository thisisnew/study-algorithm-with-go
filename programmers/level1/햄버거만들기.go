package main

import "fmt"

func main() {
	fmt.Println(햄버거만들기([]int{2, 1, 1, 2, 3, 1, 2, 3, 1}))
}

//2, 1, 1, 2, 3, 1, 2, 3, 1
//2, 1, 2, 3, 1
//2

func 햄버거만들기(ingredient []int) int {

	var result int

	for {

		hasFinishedBurger := false
		var bIdx = 0

		for i := 0; i < len(ingredient)-3; i++ {
			if isValidBurger(ingredient[i : i+4]) {
				result++
				hasFinishedBurger = true
				bIdx = i
				break
			}
		}

		ingredient = append(ingredient[:bIdx], ingredient[bIdx+4:]...)

		if !hasFinishedBurger || len(ingredient) < 4 {
			return result
		}
	}
}

func isValidBurger(burger []int) bool {
	var ex = []int{1, 2, 3, 1}

	for i, b := range burger {
		if ex[i] != b {
			return false
		}
	}

	return true
}

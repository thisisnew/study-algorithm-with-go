package main

import (
	"fmt"
	"strconv"
)

const bread = "1"
const vegetable = "2"
const meat = "3"

func main() {
	fmt.Println(햄버거만들기([]int{2, 1, 1, 2, 3, 1, 2, 3, 1}))
}

func 햄버거만들기(ingredient []int) int {

	var result = 0
	var isBurgerMade = true
	var ingredientStr = getIngredientString(ingredient)

	for {
		ln := len([]rune(ingredientStr))

		if ln < 4 || !isBurgerMade {
			break
		}

		isBurgerMade = false

		for i := 0; i < ln; i++ {
			if ingredientStr[i:i+1] != bread {
				continue
			}

			if isAnyIngredientRemain(ln, i) {
				if ingredientStr[i+1:i+2] != vegetable {
					continue
				}

				if ingredientStr[i+2:i+3] != meat {
					continue
				}

				if ingredientStr[i+3:i+4] != bread {
					continue
				}

				ingredientStr = ingredientStr[:i] + ingredientStr[i+4:]
				isBurgerMade = true
				result++
				break
			}
		}
	}

	return result
}

func getIngredientString(ingredient []int) string {

	var sum int
	var last = len(ingredient) - 1
	for i, p := range ingredient {
		sum += p

		if i < last {
			sum *= 10
		}

	}

	return strconv.Itoa(sum)
}

func isAnyIngredientRemain(ln, idx int) bool {
	if idx+3 > ln-1 {
		return false
	}

	return true
}

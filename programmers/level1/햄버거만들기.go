package main

import (
	"fmt"
	"strconv"
	"strings"
)

const bread = "1"
const vegetable = "2"
const meat = "3"

func main() {
	fmt.Println(햄버거만들기([]int{1, 3, 2, 1, 2, 1, 3, 1, 2}))
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
		var i = 0
		for i < ln {
			if ingredientStr[i:i+1] != bread {
				i++
				continue
			}

			if !isAnyIngredientRemain(ln, i) {
				i++
			}

			if ingredientStr[i+1:i+2] != vegetable {
				i++
				continue
			}

			if ingredientStr[i+2:i+3] != meat {
				i += 2
				continue
			}

			if ingredientStr[i+3:i+4] != bread {
				i += 3
				continue
			}

			a := ingredientStr[:i]
			b := ingredientStr[i+3:]
			ingredientStr = a + b
			isBurgerMade = true
			result++
			break

		}
	}

	return result
}

func getIngredientString(ingredient []int) string {

	var result strings.Builder

	for _, p := range ingredient {
		result.WriteString(strconv.Itoa(p))
	}

	return result.String()
}

func isAnyIngredientRemain(ln, idx int) bool {
	return idx+3 <= ln-1
}

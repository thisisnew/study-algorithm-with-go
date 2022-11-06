package main

import (
	"fmt"
)

const bread = 1
const vegetable = 2
const meat = 3

func main() {
	fmt.Println(햄버거만들기([]int{2, 1, 1, 2, 3, 1, 2, 3, 1}))
}

func 햄버거만들기(ingredient []int) int {

	var result = 0
	var isBurgerMade = true

	for {
		ln := len(ingredient)

		if ln < 4 || !isBurgerMade {
			break
		}

		isBurgerMade = false

		for i := 0; i < len(ingredient); i++ {
			if ingredient[i] != bread {
				continue
			}

			if isAnyIngredientRemain(ln, i) {
				if ingredient[i+1] != vegetable {
					continue
				}

				if ingredient[i+2] != meat {
					continue
				}

				if ingredient[i+3] != bread {
					continue
				}

				ingredient = append(ingredient[:i], ingredient[i+4:]...)
				isBurgerMade = true
				result++
				break
			}
		}
	}

	return result
}

func isAnyIngredientRemain(ln, idx int) bool {
	if idx+3 > ln-1 {
		return false
	}

	return true
}

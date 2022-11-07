package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(햄버거만들기([]int{2, 1, 1, 2, 3, 1, 2, 3, 1}))
}

func 햄버거만들기(ingredient []int) int {

	var result int
	var cIngredient = ingredientToString(ingredient)

	for {
		nIngredient := strings.Replace(cIngredient, "1231", "", 1)

		if nIngredient == cIngredient {
			return result
		}
		cIngredient = nIngredient
		result++
	}
}

func ingredientToString(ingredient []int) string {

	var result strings.Builder

	for _, p := range ingredient {
		result.WriteString(strconv.Itoa(p))
	}

	return result.String()

}

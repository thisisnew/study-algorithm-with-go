package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(푸드파이트대회([]int{1, 3, 4, 6}))
}

func 푸드파이트대회(food []int) string {

	var result strings.Builder

	displayOne := displayDishes(food)

	result.WriteString(displayOne)
	result.WriteString("0")
	result.WriteString(displayDishesAnother(displayOne))

	return result.String()
}

func displayDishes(food []int) string {

	var result strings.Builder

	for i := 1; i < len(food); i++ {
		cnt := food[i] / 2
		food := strconv.Itoa(i)

		for j := 0; j < cnt; j++ {
			result.WriteString(food)
		}
	}

	return result.String()
}

func displayDishesAnother(displayOne string) string {

	var result strings.Builder

	for i := len(displayOne) - 1; i >= 0; i-- {
		result.WriteString(displayOne[i : i+1])
	}

	return result.String()
}

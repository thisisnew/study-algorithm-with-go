package main

import "fmt"

func main() {
	fmt.Println(롤케이크자르기([]int{1, 2, 1, 3, 1, 4, 1, 2}))
}

func 롤케이크자르기(topping []int) int {

	if len(topping) <= 1 {
		return -1
	}

	var result = 0

	for i := 0; i < len(topping)-1; i++ {
		pre := topping[:i+1]
		post := topping[i+1:]

		if countToppings(pre) == countToppings(post) {
			result++
		}
	}

	if result > 0 {
		return result
	}

	return -1
}

func countToppings(cake []int) int {

	var p = cake[0]
	var result = 1

	for i := 1; i < len(cake); i++ {

		if p != cake[i] {
			result++
			p = cake[i]
		}

	}

	return result
}

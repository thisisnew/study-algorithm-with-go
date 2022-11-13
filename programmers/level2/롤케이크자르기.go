package main

import "fmt"

func main() {
	fmt.Println(롤케이크자르기([]int{1, 2, 1, 3, 1, 4, 1, 2}))
}

func 롤케이크자르기(topping []int) int {

	//topping = getNonOverlappingToppings(topping)

	if len(topping) <= 1 {
		return -1
	}

	var result = 0

	pre := topping[:1]
	post := topping[1:]

	for i := 1; i < len(topping)-1; i++ {

		if countToppings(pre) == countToppings(post) {
			result++
		}
	}

	return result
}

func getNonOverlappingToppings(topping []int) []int {

	var cakeMap = map[int]bool{}
	var result []int

	for _, t := range topping {

		_, ok := cakeMap[t]

		if !ok {
			cakeMap[t] = true
			result = append(result, t)
		}

	}

	return result
}

func countToppings(cake []int) int {

	var cakeMap = map[int]int{}

	for _, c := range cake {

		_, ok := cakeMap[c]

		if !ok {
			cakeMap[c]++
		}

	}

	return len(cakeMap)
}

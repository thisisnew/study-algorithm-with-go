package main

import "fmt"

func main() {
	fmt.Println(롤케이크자르기([]int{1, 2, 1, 3, 1, 4, 1, 2}))
}

func 롤케이크자르기(topping []int) int {

	var result = 0

	if len(topping) <= 1 {
		return result
	}

	pre := topping[:1]
	post := topping[1:]
	preMap := getNonOverlappingToppings(pre)
	postMap := getNonOverlappingToppings(post)

	for i := 1; i < len(topping)-1; i++ {

		if len(preMap) == len(postMap) {
			result++
		}
	}

	return result
}

func getNonOverlappingToppings(topping []int) map[int]int {

	var result = map[int]int{}

	for _, t := range topping {
		result[t]++
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

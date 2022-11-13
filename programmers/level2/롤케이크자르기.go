package main

import "fmt"

func main() {
	fmt.Println(롤케이크자르기([]int{1, 2, 3, 1, 4}))
}

func 롤케이크자르기(topping []int) int {

	var result = 0

	if len(topping) <= 1 {
		return result
	}

	pre := getNonOverlappingToppings(topping[:1])
	post := getNonOverlappingToppings(topping[1:])

	for i := 1; i < len(topping); i++ {

		t := topping[i]
		pre[t]++

		_, ok := post[t]

		if ok {
			post[t]--

			if post[t] <= 0 {
				post = removeKeyFromToppingsMap(post)
			}
		}

		if hasSameCountToppings(pre, post) {
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

func removeKeyFromToppingsMap(m map[int]int) map[int]int {

	var result = map[int]int{}

	for k, v := range m {
		if v > 0 {
			result[k] = v
		}
	}

	return result

}

func hasSameCountToppings(pre, post map[int]int) bool {
	return len(pre) == len(post)
}

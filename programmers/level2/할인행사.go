package main

import "fmt"

func main() {
	fmt.Println(
		할인행사(
			[]string{"banana", "apple", "rice", "pork", "pot"},
			[]int{3, 2, 2, 2, 1},
			[]string{"chicken", "apple", "apple", "banana", "rice", "apple", "pork", "banana", "pork", "rice", "pot", "banana", "apple", "banana"},
		),
	)
}

func 할인행사(want []string, number []int, discount []string) int {

	var result int
	wantMap := getWantMap(want, number)

	lp := len(discount) - 9
	for i := 0; i < lp; i++ {
		tenDiscount := discount[i : i+10]
		descWantMap := decreaseWandFromDiscount(cpWantMap(wantMap), tenDiscount)

		if canAllBuy(descWantMap) {
			result++
		}
	}

	return result
}

func getWantMap(want []string, number []int) map[string]int {

	var result = make(map[string]int)

	for i, w := range want {
		result[w] = number[i]
	}

	return result
}

func decreaseWandFromDiscount(wantMap map[string]int, tenDayDiscount []string) map[string]int {

	for _, d := range tenDayDiscount {
		wantMap[d]--
	}

	return wantMap
}

func canAllBuy(wantMap map[string]int) bool {

	for _, v := range wantMap {
		if v > 0 {
			return false
		}
	}

	return true
}

func cpWantMap(wantMap map[string]int) map[string]int {
	var result = make(map[string]int)

	for k, v := range wantMap {
		result[k] = v
	}

	return result
}

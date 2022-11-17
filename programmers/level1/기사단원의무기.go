package main

import "fmt"

func main() {
	fmt.Println(기사단원의무기(5, 3, 2))
}

func 기사단원의무기(number int, limit int, power int) int {
	return sumPower(getDivisorsCountSlice(number), limit, power)
}

func sumPower(divisorCount []int, limit int, power int) int {
	var result = 0

	for _, cnt := range divisorCount {
		if cnt > limit {
			result += power
			continue
		}

		result += cnt
	}

	return result
}

func getDivisorsCountSlice(number int) []int {

	var result []int

	for i := 1; i <= number; i++ {
		result = append(result, countDivisors(i))
	}

	return result
}

func countDivisors(number int) int {

	var result int

	for i := 1; i <= number; i++ {
		if number%i == 0 {
			result++
		}
	}

	return result
}

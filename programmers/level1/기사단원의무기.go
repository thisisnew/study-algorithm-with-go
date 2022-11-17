package main

import "fmt"

func main() {
	fmt.Println(기사단원의무기(5, 3, 2))
}

func 기사단원의무기(number int, limit int, power int) int {
	return sumPower(number, limit, power)
}

func sumPower(number int, limit int, power int) int {
	var result = 0

	for i := 1; i <= number; i++ {
		cnt := countDivisors(i)

		if cnt > limit {
			result += power
			continue
		}

		result += cnt
	}

	return result
}

func countDivisors(number int) int {
	var result int

	for i := 1; i*i <= number; i++ {
		if i*i == number {
			result++
			continue
		}

		if number%i == 0 {
			result += 2
			continue
		}
	}

	return result
}

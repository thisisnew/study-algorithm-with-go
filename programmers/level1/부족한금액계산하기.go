package main

import "fmt"

func main() {
	fmt.Println(부족한금액계산하기(3, 20, 4))
}

func 부족한금액계산하기(price int, money int, count int) int64 {

	cost := cost(count, price)
	result := money - cost

	if result > 0 {
		return 0
	} else {
		result = cost - money
	}

	return int64(result)
}

func cost(count, price int) int {

	var result int

	for i := 1; i <= count; i++ {
		result += i * price
	}

	return result
}

package main

import "fmt"

func main() {
	fmt.Println(치킨쿠폰(100))
}

func 치킨쿠폰(chicken int) int {

	var result int

	for {
		if chicken < 10 {
			return result
		}

		chicken -= 9
		result++
	}
}

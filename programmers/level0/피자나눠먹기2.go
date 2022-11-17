package main

import "fmt"

func main() {
	fmt.Println(피자나눠먹기2(6))
}

func 피자나눠먹기2(n int) int {

	var pizzas = 6

	for {
		if pizzas%n == 0 {
			return pizzas / 6
		}

		pizzas += 6
	}

}

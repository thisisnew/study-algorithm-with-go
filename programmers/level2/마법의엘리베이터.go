package main

import "fmt"

func main() {
	fmt.Println(마법의엘리베이터(16))
}

func 마법의엘리베이터(storey int) int {

	var result = getTotalStones(storey)

	for {

		stones := getTotalStones(storey)

		if stones < result {
			result = stones
		}

	}

	return result
}

func getTotalStones(storey int) int {

	var div = 100000000
	var result int

	for {

		if storey == 0 {
			return result
		}

		if storey/div == 0 {
			div = div / 10
			continue
		}

		result += storey / div
		storey = storey % div

	}
}

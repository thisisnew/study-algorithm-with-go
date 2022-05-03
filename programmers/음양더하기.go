package main

import "fmt"

const positiveNumber = true
const negativeNumber = false

func main() {

	absolutes := []int{1, 2, 3}
	signs := []bool{false, false, true}

	fmt.Println(음양더하기(absolutes, signs))
}

func 음양더하기(absolutes []int, signs []bool) int {

	var result int

	for i, abs := range absolutes {

		switch signs[i] {
		case positiveNumber:
			result += abs
		case negativeNumber:
			result -= abs
		}

	}

	return result

}

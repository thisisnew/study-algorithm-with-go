package main

import "fmt"

const maximumCount = 500

func main() {
	fmt.Println(콜라츠추측(6))
}

func 콜라츠추측(num int) int {

	var cnt int
	var result int

	for {

		cnt++

		if isEven(num) {
			num = num / 2
		} else {
			num = num*3 + 1
		}

		if num == 1 {
			result = cnt
			break
		}

		if cnt > maximumCount {
			result = -1
			break
		}

	}

	return result
}

func isEven(num int) bool {

	if num%2 == 0 {
		return true
	}

	return false
}

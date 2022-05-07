package main

import "fmt"

const maximumCount = 500

func main() {
	fmt.Println(콜라츠추측(7999999))
}

func 콜라츠추측(num int) int {

	var cnt int
	var result int

	for {

		if num == 1 {
			result = cnt
			break
		}

		if isEven(num) {
			num = num / 2
		} else {
			num = num*3 + 1
		}

		cnt++

		if cnt > maximumCount {
			result = -1
			break
		}
	}

	return result
}

func isEven(num int) bool {

	switch num % 2 {
	case 0:
		return true
	default:
		return false
	}
}

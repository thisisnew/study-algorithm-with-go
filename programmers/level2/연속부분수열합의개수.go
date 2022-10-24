package main

import "fmt"

func main() {
	fmt.Println(연속부분수열합의개수([]int{7, 9, 1, 1, 4}))
}

func 연속부분수열합의개수(elements []int) int {

	var result int
	var dif = 1

	for {

		if dif == len(elements) {
			break
		}

		dif++
	}

	return result
}

func sumElementByDif(elements []int, dif int) int {
	var result int
	var cnt = 0
	var idx = 0

	for {
		result += elements[idx]
		idx++
		cnt++

		if cnt == dif {
			break
		}

		if idx == len(elements)-1 {
			idx = 0
		}
	}

	return result
}

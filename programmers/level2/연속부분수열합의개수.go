package main

import "fmt"

func main() {
	fmt.Println(연속부분수열합의개수([]int{7, 9, 1, 1, 4}))
}

func 연속부분수열합의개수(elements []int) int {

	var result int
	var cnt = 1
	var maxCnt = len(elements)

	for {

		if cnt == maxCnt {
			break
		}

		cnt++
	}

	return result
}

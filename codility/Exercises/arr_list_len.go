package main

import "fmt"

func main() {
	fmt.Println(ArrListLen([]int{1, 4, -1, 3, 2}))
}

func ArrListLen(A []int) int {

	var result = 0
	var idx = 0

	for {
		if A[idx] == -1 {
			return result
		}

		if A[idx] > result {
			result = A[idx]
		}

		idx = A[idx]
	}

}

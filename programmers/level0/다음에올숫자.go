package main

import "fmt"

func main() {
	fmt.Println(다음에올숫자([]int{2, 4, 8}))
}

func 다음에올숫자(common []int) int {

	for i := 0; i < len(common)-1; i++ {

		if common[i] == 0 || common[i+1] == 0 {
			return common[len(common)-1] + (common[len(common)-1] - common[len(common)-2])
		}

		if common[i+1]%common[i] != 0 { //등차
			return common[len(common)-1] + (common[len(common)-1] - common[len(common)-2])
		}
	}

	return common[len(common)-1] * (common[1] / common[0])
}

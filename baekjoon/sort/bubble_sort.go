package main

import "fmt"

func main() {

	var items = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

	for i := len(items); i > 0; i-- {
		isSwap := false

		for j := 0; j < i-1; j++ {
			if items[j] > items[j+1] {
				items[j+1], items[j] = items[j], items[j+1]
				isSwap = true
			}
		}

		if !isSwap {
			break
		}
	}

	fmt.Println(items)

}

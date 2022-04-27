package main

import "fmt"

func main() {

	var items = []int{7, 5, 9, 0, 1, 6, 2, 4, 8, 3}

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

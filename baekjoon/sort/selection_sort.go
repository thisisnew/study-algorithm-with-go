package main

import "fmt"

func main() {

	var items = []int{7, 5, 9, 0, 1, 6, 2, 4, 8, 3}

	for i := 0; i < len(items); i++ {
		var minIndex = i

		for j := i; j < len(items); j++ {
			if items[j] < items[minIndex] {
				minIndex = j
			}
		}

		items[i], items[minIndex] = items[minIndex], items[i]
	}

	fmt.Println(items)
}

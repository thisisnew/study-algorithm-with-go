package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var items = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	quickSort(items)
	fmt.Println(items)
}

func quickSort(items []int) []int {

	if len(items) < 2 {
		return items
	}

	left := 0
	right := len(items) - 1
	pivot := rand.Int() % len(items)

	items[pivot], items[right] = items[right], items[pivot]

	for i, _ := range items {
		if items[i] < items[right] {
			items[left], items[i] = items[i], items[left]
			left++
		}
	}

	quickSort(items[:left])
	quickSort(items[left+1:])

	return items
}

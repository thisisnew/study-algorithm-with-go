package main

import "fmt"

func main() {

	var items = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

	for i := 1; i < len(items); i++ {
		for j := i; j > 0; j-- {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
		}
	}

	fmt.Println(items)
}

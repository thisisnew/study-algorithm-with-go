package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var n int
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &n)

	var arr []int

	for i := 0; i < n; i++ {
		var num int
		fmt.Fscanln(read, &num)
		arr = append(arr, num)
	}

	for i := 0; i < len(arr); i++ {
		var minIndex = i

		for j := i; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}

	min := arr[0]
	for _, item := range arr {
		if item > min {
			fmt.Print(min)
			min = item
		}
	}

}

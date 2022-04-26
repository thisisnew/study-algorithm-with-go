package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var n int
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &n)

	text, _, _ := read.ReadLine()

	nums := strings.Split(string(text), " ")

	var arr []int

	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(nums[i])
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

	var min int

	for i, item := range arr {

		if i == 0 {
			fmt.Printf("%v ", item)
			min = item
			continue
		}

		if item > min {
			fmt.Printf("%v ", item)
			min = item
		}
	}

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {

	var read = bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(read, &n)

	var nums = make([]string, n)

	input, _, _ := read.ReadLine()
	inputs := strings.Split(string(input), " ")

	for i, n := range inputs {
		nums[i] = n
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

out:
	for i, n := range nums {

		for j, prop := range inputs {

			if n == prop {
				fmt.Print(j + 1)
				continue out
			}

		}

		if i < len(nums)-1 {
			fmt.Print(" ")
		}
	}

}

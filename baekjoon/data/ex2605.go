package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var read = bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(read, &n)

	var inputs = map[int]int{}
	var nums = make([]int, n)

	input, _, _ := read.ReadLine()

	for i, in := range strings.Split(string(input), " ") {
		n, _ := strconv.Atoi(in)
		inputs[n] = i + 1
		nums[i] = n
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	for i, n := range nums {
		fmt.Print(inputs[n])
		if i < len(nums)-1 {
			fmt.Print(" ")
		}
	}

}

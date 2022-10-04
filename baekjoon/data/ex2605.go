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

	var m = map[int]int{}
	var nums = make([]int, n)

	input, _, _ := read.ReadLine()
	inputs := strings.Split(string(input), " ")

	for i, in := range inputs {
		n, _ := strconv.Atoi(in)
		m[n] = i + 1
		nums[i] = n
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	for i, n := range nums {
		fmt.Print(m[n])
		if i < len(nums)-1 {
			fmt.Print(" ")
		}
	}

}

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
	var n int
	var read = bufio.NewReader(os.Stdin)

	fmt.Fscanln(read, &n)

	var sliceA = make([]int, n)
	text, _, _ := read.ReadLine()

	for i, t := range strings.Split(string(text), " ") {
		num, _ := strconv.Atoi(t)
		sliceA[i] = num
	}

	sort.Slice(sliceA, func(i, j int) bool {
		return sliceA[i] < sliceA[j]
	})

	var m int
	fmt.Fscanln(read, &m)
	text, _, _ = read.ReadLine()
	for _, t := range strings.Split(string(text), " ") {
		num, _ := strconv.Atoi(t)

		if getResultByBinarySearch(sliceA, num) < 0 {
			fmt.Println(0)
		} else {
			fmt.Println(1)
		}
	}
}

func getResultByBinarySearch(sliceA []int, num int) int {
	var start = 0
	var end = len(sliceA) - 1

	for start <= end {

		var middle = (start + end) / 2

		if num < sliceA[middle] {
			end = middle - 1
		} else if num > sliceA[middle] {
			start = middle + 1
		} else {
			return middle
		}

	}

	return -1
}

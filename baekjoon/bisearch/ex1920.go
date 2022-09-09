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

	inputSl := strings.Split(string(text), " ")

	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(inputSl[i])
		sliceA[i] = num
	}

	sort.Slice(sliceA, func(i, j int) bool {
		return sliceA[i] < sliceA[j]
	})

	var m int
	fmt.Fscanln(read, &m)
	text, _, _ = read.ReadLine()
	sliceM := strings.Split(string(text), " ")

	for i := 0; i < m; i++ {
		num, _ := strconv.Atoi(sliceM[i])

		if getResultByBinarySearch(sliceA, num) >= 0 {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	}
}

func getResultByBinarySearch(sliceA []int, num int) int {
	var start = 0
	var end = len(sliceA) - 1
	var middle = 0

	for start <= end {

		middle = (start + end) / 2

		if middle >= len(sliceA) || middle < 0 {
			return -1
		}

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

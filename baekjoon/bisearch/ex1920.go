package main

import (
	"bufio"
	"fmt"
	"math"
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
		fmt.Println(getResultByBinarySearch(sliceA, num))
	}
}

func getResultByBinarySearch(sliceA []int, num int) int {
	var start = 0
	var end = len(sliceA)
	var middle = int(math.Floor(float64(start + end/2)))

	for sliceA[middle] != num && start <= end {

		if num > sliceA[middle] {
			start = middle + 1
		} else {
			end = middle - 1
		}

		middle = int(math.Floor(float64(start+end) / 2))
	}

	if sliceA[middle] == num {
		return 1
	}

	return 0
}

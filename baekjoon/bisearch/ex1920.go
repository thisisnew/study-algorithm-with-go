package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	var read = bufio.NewReader(os.Stdin)

	fmt.Fscanln(read, &n)

	var sliceA = make([]int32, n)
	text, _, _ := read.ReadLine()

	for i, t := range strings.Split(string(text), " ") {
		num, _ := strconv.ParseInt(t, 10, 32)
		sliceA[i] = int32(num)
	}

	var m int
	fmt.Fscanln(read, &m)
	text, _, _ = read.ReadLine()
	for _, t := range strings.Split(string(text), " ") {
		num, _ := strconv.ParseInt(t, 10, 32)
		fmt.Println(getResultByBinarySearch(sliceA, int32(num)))
	}
}

func getResultByBinarySearch(sliceA []int32, num int32) int {
	var start int32
	var end = int32(len(sliceA))
	var middle = int32(math.Floor(float64(start + end/2)))

	for middle != num && start <= end {

		if num > middle {
			start = middle + 1
		} else {
			end = middle - 1
		}

		middle = int32(math.Floor(float64(start+end) / 2))
	}

	if middle == num {
		return 1
	}

	return 0
}

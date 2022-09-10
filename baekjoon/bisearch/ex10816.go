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

	var sl = make([]int, n)
	text, _, _ := read.ReadLine()

	inputSl := strings.Split(string(text), " ")

	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(inputSl[i])
		sl[i] = num
	}

	sort.Slice(sl, func(i, j int) bool {
		return sl[i] < sl[j]
	})

	var m int
	fmt.Fscanln(read, &m)
	text, _, _ = read.ReadLine()
	sliceM := strings.Split(string(text), " ")

	var result = make([]int, m)

	for i := 0; i < m; i++ {
		num, _ := strconv.Atoi(sliceM[i])
		result[i] = countNumbersUseBinarySearch(sl, num)
	}

	for i, n := range result {
		fmt.Print(n)

		if i != len(result)-1 {
			fmt.Print(" ")
		}
	}

}

func countNumbersUseBinarySearch(sl []int, num int) int {

	var result = 0
	var start = 0
	var end = len(sl)
	var middle = int(math.Floor(float64((start + end) / 2)))
	var isResultAdd = false

	for start <= end {

		if middle >= len(sl) {
			return result
		}

		if sl[middle] == num && !isResultAdd {
			result++
			middle++
			isResultAdd = true
		} else {
			if sl[middle] < num {
				start = middle + 1
			}

			if sl[middle] > num {
				end = middle - 1
			}

			m := int(math.Floor(float64((start + end) / 2)))

			if middle == m {
				middle++
			} else {
				middle = m
			}

			isResultAdd = false
		}
	}

	return result
}

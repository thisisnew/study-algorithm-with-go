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

	text, _, _ := read.ReadLine()
	inputSl := strings.Split(string(text), " ")

	var sl = make([]int, len(inputSl))
	for i := 0; i < len(inputSl); i++ {
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

	var result = make([]int, len(sliceM))

	for i := 0; i < len(sliceM); i++ {
		num, _ := strconv.Atoi(sliceM[i])
		result[i] = checkExistByBinarySearch(sl, num)
	}

	for i, n := range result {
		fmt.Print(n)

		if i != len(result)-1 {
			fmt.Print(" ")
		}
	}

}

func checkExistByBinarySearch(sl []int, num int) int {

	var start = 0
	var end = len(sl) - 1
	var middle = int(math.Floor(float64((start + end) / 2)))

	for start <= end {

		if middle >= len(sl) {
			return 0
		}

		if sl[middle] < num {
			start = middle + 1
		}

		if sl[middle] > num {
			end = middle - 1
		}

		middle = int(math.Floor(float64((start + end) / 2)))

		if sl[middle] == num {
			return 1
		}
	}

	if middle < len(sl) && sl[middle] == num {
		return 1
	}

	return 0
}

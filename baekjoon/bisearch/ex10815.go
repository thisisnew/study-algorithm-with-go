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

	text, _ := read.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	inputSl := strings.Split(text, " ")

	var sl = make([]int, n)
	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(inputSl[i])
		sl[i] = num
	}

	sort.Slice(sl, func(i, j int) bool {
		return sl[i] < sl[j]
	})

	var m int
	fmt.Fscanln(read, &m)
	text, _ = read.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	sliceM := strings.Split(text, " ")

	for i := 0; i < m; i++ {
		num, _ := strconv.Atoi(sliceM[i])
		fmt.Print(checkExistByBinarySearch(sl, num))

		if i < len(sliceM)-1 {
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

		if middle < len(sl) && sl[middle] == num {
			return 1
		}
	}

	if middle < len(sl) && sl[middle] == num {
		return 1
	}

	return 0
}

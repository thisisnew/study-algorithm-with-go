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
	text, _ := read.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	inputSl := strings.Split(text, " ")

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

	var result []int
	var start = 0
	var end = len(sl)
	var middle = int(math.Floor(float64((start + end) / 2)))

	for start <= end {

		if middle >= len(sl) {
			return len(result)
		}

		if sl[middle] == num {

			result = append(result, num)

			var minusMiddle = middle - 1
			var plusMiddle = middle + 1

			for {

				if minusMiddle >= 0 {
					if sl[minusMiddle] == num {
						result = append(result, num)
					}
				}

				minusMiddle--

				if plusMiddle < len(sl) {
					if sl[plusMiddle] == num {
						result = append(result, num)
					}
				}

				plusMiddle++

				if minusMiddle < 0 && plusMiddle >= len(sl) {
					return len(result)
				}
			}

		}

		if sl[middle] < num {
			start = middle + 1
		}

		if sl[middle] > num {
			end = middle - 1
		}

		middle = int(math.Floor(float64((start + end) / 2)))
	}

	return len(result)
}

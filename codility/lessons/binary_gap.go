package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(binaryGap(1041))
}

func binaryGap(N int) int {

	var binNum = strconv.FormatInt(int64(N), 2)
	var result = 0
	var zeros = 0
	var isStart = false

	for _, n := range binNum {

		if string(n) == "1" {
			if !isStart {
				isStart = true
				continue
			}

			isStart = false

			if zeros > result {
				result = zeros
			}
		}

		if isStart {
			zeros++
		}
	}

	return result
}

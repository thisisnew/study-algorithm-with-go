package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(binaryGap(1162))
}

func binaryGap(N int) int {

	var binNum = strconv.FormatInt(int64(N), 2)
	var idx = strings.Index(binNum, "1")
	var zeros = 0
	var result = 0

	for {
		idx++

		if idx > len([]rune(binNum))-1 {
			return result
		}

		n := binNum[idx : idx+1]

		switch n {
		case "1":
			if zeros > result {
				result = zeros
			}

			zeros = 0
		default:
			zeros++
		}
	}
}

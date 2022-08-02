package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {

	var n int
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &n)

	for i := 0; i < n; i++ {
		var input string
		fmt.Fscanln(read, &input)

		var sq = math.Sqrt(float64(len([]rune(input))))
		var idx = int(sq) - 1
		var tempIdx = idx
		var tempCnt = 1
		var totalCnt = 0
		var result strings.Builder

		for {
			if totalCnt == len([]rune(input)) {
				break
			}

			if tempCnt > int(sq) {
				tempIdx = idx - 1
				idx = tempIdx
				tempCnt = 1
			}

			result.WriteString(input[tempIdx : tempIdx+1])
			tempIdx += int(sq)
			tempCnt++
			totalCnt++
		}

		fmt.Println(result.String())
	}

}

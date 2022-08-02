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

		var sl = make([][]string, int(sq), int(sq))
		var idx = 0

		for j := 0; j < len(sl); j++ {
			sl[j] = make([]string, int(sq))

			for k := 0; k < len(sl[j]); k++ {
				sl[j][k] = input[idx : idx+1]
				idx++
			}
		}

		var result strings.Builder

		for j := len(sl) - 1; j >= 0; j-- {
			for k := 0; k < len(sl); k++ {
				s := sl[k][j]
				result.WriteString(s)
			}
		}

		fmt.Println(result.String())
	}

}

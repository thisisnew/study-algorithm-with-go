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

	for i := 0; i < n; i++ {
		var input string
		fmt.Fscanln(read, &input)

		words := strings.Split(input, "-")

		var pre = calculatePreWordValue(0, 0, float64(len(words[0])))
		var post, _ = strconv.Atoi(words[1])

		if math.Abs(float64(pre-post)) <= 100 {
			fmt.Println("nice")
		} else {
			fmt.Println("not nice")
		}

	}
}

func calculatePreWordValue(total int, idx, len float64) int {

	for {
		if len < 0 {
			return total
		}

		total += int(power(len-1) * idx)
		idx++
		len--
	}
}

func power(len float64) float64 {
	return math.Pow(26, len)
}

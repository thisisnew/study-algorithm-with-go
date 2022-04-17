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

		var pre int
		var power = float64(len(words[0]) - 1)

		for j := 0; j < len(words[0]); j++ {
			pre += j * int(math.Pow(26, power))
			power--
		}

		var post, _ = strconv.Atoi(words[1])

		if math.Abs(float64(pre-post)) <= 100 {
			fmt.Println("nice")
		} else {
			fmt.Println("not nice")
		}

	}
}

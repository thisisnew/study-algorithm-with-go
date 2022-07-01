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

	text, _, _ := read.ReadLine()

	sl := strings.Split(string(text), " ")

	var arr []string
	var dp = make([]float64, n)

	var result float64
	for i, s := range sl {
		arr = append(arr, s)
		dp[i] = 1

		for j := i - 1; j >= 0; j-- {
			if arr[i] < arr[j] && dp[i] <= dp[j] {
				dp[i] = dp[j] + 1
			}
		}

		result = math.Max(result, dp[i])
	}

	fmt.Println(result)
}

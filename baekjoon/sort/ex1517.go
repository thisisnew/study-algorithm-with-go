package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &n)

	text, _, _ := read.ReadLine()
	nums := strings.Split(string(text), " ")

	var result int

	for i := n; i > 0; i-- {
		isSwap := false

		for j := 0; j < i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
				isSwap = true
				result++
			}
		}

		if !isSwap {
			break
		}
	}

	fmt.Println(result)
}

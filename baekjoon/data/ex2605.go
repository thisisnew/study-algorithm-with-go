package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var read = bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(read, &n)

	input, _, _ := read.ReadLine()
	inputs := strings.Split(string(input), " ")

	var result []int
	for i, num := range inputs {
		n, _ = strconv.Atoi(num)
		number := i + 1

		switch {
		case i == 0:
			result = []int{number}
		case n == 0:
			result = append(result, number)
		case n >= len(result):
			result = append([]int{number}, result...)
		default:
			index := i - n
			result = append(result[:index+1], result[index:]...)
			result[index] = number
		}
	}

	for i, n := range result {
		fmt.Print(n)
		if i < len(result)-1 {
			fmt.Print(" ")
		}
	}
}

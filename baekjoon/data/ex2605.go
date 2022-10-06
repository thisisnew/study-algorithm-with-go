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
		case n == 0:
			result = append(result, number)
		case n >= len(result):
			result = append([]int{number}, result...)
		case n < len(result):
			index := i - n
			result = append(result[:index+1], result[index:]...)
			result[index] = number
		}
	}

	fmt.Println(strings.Trim(fmt.Sprint(result), "[]"))
}

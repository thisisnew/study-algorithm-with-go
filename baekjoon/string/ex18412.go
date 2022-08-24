package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var dictionary map[string]string

func main() {

	var read = bufio.NewReader(os.Stdin)
	text, _, _ := read.ReadLine()
	nums := strings.Split(string(text), " ")

	var n, a, b int

	n, _ = strconv.Atoi(nums[0])
	a, _ = strconv.Atoi(nums[1])
	b, _ = strconv.Atoi(nums[2])

	var s string
	fmt.Fscanln(read, &s)

	if a == b {
		fmt.Println(s)
		return
	}

	var result = make([]string, n)

	for i := 0; i < a-1; i++ {
		result[i] = s[i : i+1]
	}

	var idx = a - 1
	for i := b; i >= a; i-- {
		result[idx] = s[i-1 : i]
		idx++
	}

	for i := b; i < n; i++ {
		result[i] = s[i : i+1]
	}

	fmt.Println(strings.Join(result, ""))
}

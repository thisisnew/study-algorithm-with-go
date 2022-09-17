package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var read = bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(read, &n)

	text, _, _ := read.ReadLine()

	var sl = make([]int, n)

	for i, s := range strings.Split(string(text), " ") {
		num, _ := strconv.Atoi(s)
		sl[i] = num
	}

	sort.Slice(sl, func(i, j int) bool {
		return sl[i] < sl[j]
	})

	fmt.Println(sl[len(sl)-1] - sl[0])
}

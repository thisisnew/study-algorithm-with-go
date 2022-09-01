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
	var t int
	fmt.Fscanln(read, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscanln(read, &n)
		fmt.Println(getMaximumDrunkSchool(read, n))
	}

}

func getMaximumDrunkSchool(read *bufio.Reader, n int) string {

	var alcohols = make([]int, n)
	var schools = map[int]string{}

	for i := 0; i < n; i++ {
		text, _, _ := read.ReadLine()

		sl := strings.Split(string(text), " ")
		alc, _ := strconv.Atoi(sl[1])

		alcohols[i] = alc
		schools[alc] = sl[0]
	}

	sort.Slice(alcohols, func(i, j int) bool {
		return alcohols[i] > alcohols[j]
	})

	return schools[alcohols[0]]
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type stringArray []string

func main() {
	var input string
	var reader = bufio.NewReader(os.Stdin)

	fmt.Fscan(reader, &input)

	array := stringToStringArray(input)

	sort.Sort(stringArray(array))

	fmt.Println(array)
}

func stringToStringArray(input string) []string {
	return strings.Split(input, "")
}

func (s stringArray) Len() int {
	return len(s)
}
func (s stringArray) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s stringArray) Less(i, j int) bool {
	si, _ := strconv.Atoi(s[i])
	sj, _ := strconv.Atoi(s[j])

	return si > sj
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var n int
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &n)

	fmt.Println(strings.Index(getTotalNumberString(n), strconv.Itoa(n)) + 1)
}

func getTotalNumberString(n int) string {

	var result strings.Builder

	for i := 1; i <= n; i++ {
		d := strconv.Itoa(i)
		result.WriteString(d)
	}

	return result.String()
}

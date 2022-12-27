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

	var sb strings.Builder

	for i := 1; i <= 100000; i++ {
		if i > n {
			break
		}

		sb.WriteString(strconv.Itoa(i))
	}

}

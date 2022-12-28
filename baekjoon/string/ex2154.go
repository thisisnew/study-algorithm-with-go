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
	var div = 10

	for i := 1; i <= 100000; i++ {

		if i == div*10 {
			div *= 10
		}

		d := strconv.Itoa(i / div)

		if d == strconv.Itoa(n)[0:1] {

		}

		sb.WriteString(strconv.Itoa(i))
	}

}

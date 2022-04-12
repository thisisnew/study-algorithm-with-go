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
	var m int
	var read = bufio.NewReader(os.Stdin)

	fmt.Fscanln(read, &n, &m)

	w := strconv.Itoa(n)

	var b strings.Builder

	for i := 0; i < n; i++ {
		b.WriteString(w)
	}

	result := b.String()

	if len(result) > m {
		fmt.Println(result[:m])
	} else {
		fmt.Println(result)
	}

}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {

	var n int
	var read = bufio.NewReader(os.Stdin)
	var b bytes.Buffer

	fmt.Fscanln(read, &n)

	for i := 0; i < n; i++ {
		text, _, _ := read.ReadLine()
		input := string(text)

		b.WriteString(strings.ToUpper(input[0:1]))
		b.WriteString(input[1:])
		fmt.Println(b.String())
		b.Reset()
	}
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {

	var t int
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &t)

	for i := 0; i < t; i++ {
		var pos int
		var word string
		fmt.Fscanln(read, &pos, &word)

		var b bytes.Buffer
		b.WriteString(word[0 : pos-1])
		b.WriteString(word[pos:])

		fmt.Println(b.String())
	}
}

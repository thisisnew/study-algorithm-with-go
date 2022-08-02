package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func main() {

	var input string
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &input)

	h := md5.New()
	io.WriteString(h, input)
	fmt.Printf("%x", h.Sum(nil))
}

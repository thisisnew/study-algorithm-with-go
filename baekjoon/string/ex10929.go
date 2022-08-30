package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {

	var read = bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscanln(read, &s)

	sum := sha256.Sum224([]byte(s))
	fmt.Printf("%x", sum)
}
